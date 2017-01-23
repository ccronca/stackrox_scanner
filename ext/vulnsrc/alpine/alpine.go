// Copyright 2017 clair authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package alpine implements a vulnerability source updater using the
// alpine-secdb git repository.
package alpine

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"gopkg.in/yaml.v2"

	"github.com/coreos/pkg/capnslog"

	"github.com/coreos/clair"
	"github.com/coreos/clair/database"
	"github.com/coreos/clair/ext/versionfmt"
	"github.com/coreos/clair/ext/versionfmt/dpkg"
	"github.com/coreos/clair/ext/vulnsrc"
	"github.com/coreos/clair/pkg/commonerr"
)

const (
	// When available, this should be updated to use HTTPS.
	secdbGitURL  = "http://git.alpinelinux.org/cgit/alpine-secdb"
	updaterFlag  = "alpine-secdbUpdater"
	nvdURLPrefix = "https://cve.mitre.org/cgi-bin/cvename.cgi?name="
)

var (
	log = capnslog.NewPackageLogger("github.com/coreos/clair", "ext/vulnsrc/alpine")
)

func init() {
	vulnsrc.RegisterUpdater("alpine", &updater{})
}

type updater struct {
	repositoryLocalPath string
}

func (u *updater) Update(db database.Datastore) (resp vulnsrc.UpdateResponse, err error) {
	log.Info("fetching Alpine vulnerabilities")

	// Pull the master branch.
	var commit string
	commit, err = u.pullRepository()
	if err != nil {
		return
	}

	// Ask the database for the latest commit we successfully applied.
	var dbCommit string
	dbCommit, err = db.GetKeyValue(updaterFlag)
	if err != nil {
		return
	}

	// Set the updaterFlag to equal the commit processed.
	resp.FlagName = updaterFlag
	resp.FlagValue = commit

	// Short-circuit if there have been no updates.
	if commit == dbCommit {
		log.Debug("no alpine update")
		return
	}

	var namespaces []string
	namespaces, err = detectNamespaces(u.repositoryLocalPath)
	// Append any changed vulnerabilities to the response.
	for _, namespace := range namespaces {
		var vulns []database.Vulnerability
		var note string
		vulns, note, err = parseVulnsFromNamespace(u.repositoryLocalPath, namespace)
		if err != nil {
			return
		}
		if note != "" {
			resp.Notes = append(resp.Notes, note)
		}
		resp.Vulnerabilities = append(resp.Vulnerabilities, vulns...)
	}

	return
}

func (u *updater) Clean() {
	if u.repositoryLocalPath != "" {
		os.RemoveAll(u.repositoryLocalPath)
	}
}

func detectNamespaces(path string) ([]string, error) {
	// Open the root directory.
	dir, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer dir.Close()

	// Get a list of the namspaces from the directory names.
	finfos, err := dir.Readdir(0)
	if err != nil {
		return nil, err
	}

	var namespaces []string
	for _, info := range finfos {
		if !info.IsDir() {
			continue
		}
		// Filter out hidden directories like `.git`.
		if strings.HasPrefix(info.Name(), ".") {
			continue
		}

		namespaces = append(namespaces, info.Name())
	}

	return namespaces, nil
}

type parserFunc func(io.Reader) ([]database.Vulnerability, error)

var parsers = map[string]parserFunc{
	"v3.3": parse33YAML,
	"v3.4": parse34YAML,
}

func parseVulnsFromNamespace(repositoryPath, namespace string) (vulns []database.Vulnerability, note string, err error) {
	var file io.ReadCloser
	file, err = os.Open(repositoryPath + "/" + namespace + "/main.yaml")
	if err != nil {
		return
	}
	defer file.Close()

	parseFunc, exists := parsers[namespace]
	if !exists {
		note = fmt.Sprintf("The file %s is not mapped to any Alpine version number", namespace)
		return
	}

	vulns, err = parseFunc(file)
	return
}

func (u *updater) pullRepository() (commit string, err error) {
	// If the repository doesn't exist, clone it.
	if _, pathExists := os.Stat(u.repositoryLocalPath); u.repositoryLocalPath == "" || os.IsNotExist(pathExists) {
		if u.repositoryLocalPath, err = ioutil.TempDir(os.TempDir(), "alpine-secdb"); err != nil {
			return "", vulnsrc.ErrFilesystem
		}

		cmd := exec.Command("git", "clone", secdbGitURL, ".")
		cmd.Dir = u.repositoryLocalPath
		if out, err := cmd.CombinedOutput(); err != nil {
			u.Clean()
			log.Errorf("could not pull alpine-secdb repository: %s. output: %s", err, out)
			return "", commonerr.ErrCouldNotDownload
		}
	} else {
		// The repository already exists and it needs to be refreshed via a pull.
		cmd := exec.Command("git", "pull")
		cmd.Dir = u.repositoryLocalPath
		if _, err := cmd.CombinedOutput(); err != nil {
			return "", vulnsrc.ErrGitFailure
		}
	}

	cmd := exec.Command("git", "rev-parse", "HEAD")
	cmd.Dir = u.repositoryLocalPath
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", vulnsrc.ErrGitFailure
	}

	commit = strings.TrimSpace(string(out))
	return
}

type secdb33File struct {
	Distro   string `yaml:"distroversion"`
	Packages []struct {
		Pkg struct {
			Name    string   `yaml:"name"`
			Version string   `yaml:"ver"`
			Fixes   []string `yaml:"fixes"`
		} `yaml:"pkg"`
	} `yaml:"packages"`
}

func parse33YAML(r io.Reader) (vulns []database.Vulnerability, err error) {
	var rBytes []byte
	rBytes, err = ioutil.ReadAll(r)
	if err != nil {
		return
	}

	var file secdb33File
	err = yaml.Unmarshal(rBytes, &file)
	if err != nil {
		return
	}
	for _, pack := range file.Packages {
		pkg := pack.Pkg
		for _, fix := range pkg.Fixes {
			err = versionfmt.Valid(dpkg.ParserName, pkg.Version)
			if err != nil {
				log.Warningf("could not parse package version '%s': %s. skipping", pkg.Version, err.Error())
				continue
			}

			vulns = append(vulns, database.Vulnerability{
				Name:     fix,
				Severity: clair.Unknown,
				Link:     nvdURLPrefix + fix,
				FixedIn: []database.FeatureVersion{
					{
						Feature: database.Feature{
							Namespace: database.Namespace{
								Name:          "alpine:" + file.Distro,
								VersionFormat: dpkg.ParserName,
							},
							Name: pkg.Name,
						},
						Version: pkg.Version,
					},
				},
			})
		}
	}
	return
}

type secdb34File struct {
	Distro   string `yaml:"distroversion"`
	Packages []struct {
		Pkg struct {
			Name  string              `yaml:"name"`
			Fixes map[string][]string `yaml:"secfixes"`
		} `yaml:"pkg"`
	} `yaml:"packages"`
}

func parse34YAML(r io.Reader) (vulns []database.Vulnerability, err error) {
	var rBytes []byte
	rBytes, err = ioutil.ReadAll(r)
	if err != nil {
		return
	}

	var file secdb34File
	err = yaml.Unmarshal(rBytes, &file)
	if err != nil {
		return
	}

	for _, pack := range file.Packages {
		pkg := pack.Pkg
		for version, vulnStrs := range pkg.Fixes {
			err := versionfmt.Valid(dpkg.ParserName, version)
			if err != nil {
				log.Warningf("could not parse package version '%s': %s. skipping", version, err.Error())
				continue
			}

			for _, vulnStr := range vulnStrs {
				var vuln database.Vulnerability
				vuln.Severity = clair.Unknown
				vuln.Name = vulnStr
				vuln.Link = nvdURLPrefix + vulnStr
				vuln.FixedIn = []database.FeatureVersion{
					{
						Feature: database.Feature{
							Namespace: database.Namespace{
								Name:          "alpine:" + file.Distro,
								VersionFormat: dpkg.ParserName,
							},
							Name: pkg.Name,
						},
						Version: version,
					},
				}
				vulns = append(vulns, vuln)
			}
		}
	}

	return
}
