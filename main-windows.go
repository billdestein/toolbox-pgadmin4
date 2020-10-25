package main

//--------------------------------------------------------------------------------------------------
//
// This program will
//
//     - Compile the Python code using qmake.
//     - Compile the web assets using yarn.
//     - Create a tarball and put in ~/.toolbox-tarballs
//
// The toolbox build process will copy the tarball from ~/.toolbox-tarballs into the release
// tarball.
//
// This program also assumes that it is running in a Python >= 3.4 virtual environment.
//
// These directories have been added to the PATH environment variable.
//
// - In order to find the command line compiler 'cl.exe'
//     - C:\Program Files (x86)\Microsoft Visual Studio\2017\BuildTools\VC\Tools\MSVC\14.16.27023\bin\Hostx64\x64
//
// - In order to find 'qmake.exe'
//     - C:\Qt\Qt5.11.3\5.11.3\msvc2017_64\bin
//
//--------------------------------------------------------------------------------------------------

import (
  "fmt"
  "os"
  "os/exec"
  "path"
  "path/filepath"
)
//  "os/user"

type Builder struct {
}

func (this Builder) build() {
  repoDir, _ :=  filepath.Abs(filepath.Dir(os.Args[0]))
  runtimeDir := path.Join(repoDir, "runtime")
  // currentUser, _ := user.Current();
  // homeDir := currentUser.HomeDir    
  // tarballDir := path.Join(homeDir, "toolbox-tarballs")
  // tarballFilepath := path.Join(tarballDir, "toolbox-pgadmin4.tgz")

  // Find the qmake executable
  qmakeFilepath, err := exec.LookPath("qmake")
  if err != nil {
    fmt.Printf("Could not find qmake executable. %s\n", err.Error())
    os.Exit(1);
  }
  fmt.Printf("using qmake: '%s'\n", qmakeFilepath);

  // Find the python executable
  pythonFilepath, err := exec.LookPath("python")
  if err != nil {
    fmt.Printf("Could not find python executable. %s\n", err.Error())
    os.Exit(1);
  }
  fmt.Printf("using python: '%s'\n", pythonFilepath);

  // Find the yarn executable
  yarnFilepath, err := exec.LookPath("yarn")
  if err != nil {
    fmt.Printf("Could not find yarn executable. %s\n", err.Error())
    os.Exit(1);
  }
  fmt.Printf("using yarn: '%s'\n", yarnFilepath);

  // Go to the runtime directory
  os.Chdir(runtimeDir)
  
  // The PgAdmin process needs to know where to find Python
  pythonDir := filepath.Dir(pythonFilepath)
  fmt.Printf("setting PGADMIN_PYTHON_DIR to %s\n", pythonDir)
  os.Setenv("PGADMIN_PYTHON_DIR", pythonDir)

  // qmake
  command := exec.Command(qmakeFilepath)
  fmt.Printf("command: %s\n", command.String());
  output, err := command.CombinedOutput();
  if err != nil {
    fmt.Printf("error running command '%s' '%s'. %s\n", command.String(), output, err.Error())
    os.Exit(1)
  }

  // Exec 'qmake'
  // command = fmt.Sprintf("%s", qmakeFilepath)
  // fmt.Printf("command: '%s'\n", command)
  // os.Chdir(runtimeDir)
  // output, err = exec.Command("bash", "-c", command).Output()
  // os.Chdir(repoDir)
  // if err != nil {
  //   fmt.Printf("Failed to execute command '%s'. %s %s\n", command, output, err.Error())
  //   os.Exit(1)
  // }

  // // Exec 'qmake CONFIG+=release'
  // command = fmt.Sprintf("%s CONFIG+=release", qmakeFilepath)
  // fmt.Printf("command: '%s'\n", command)
  // os.Chdir(runtimeDir)
  // output, err = exec.Command("bash", "-c", command).Output()
  // os.Chdir(repoDir)
  // if err != nil {
  //   fmt.Printf("Failed to execute command '%s'. %s %s\n", command, output, err.Error())
  //   os.Exit(1)
  // }

  // // Build the web assets
  // webDir := path.Join(repoDir, "web")

  // // Exec 'yarn install'
  // command = fmt.Sprintf("yarn install")
  // fmt.Printf("command: '%s'\n", command)
  // os.Chdir(webDir)
  // output, err = exec.Command("bash", "-c", command).Output()
  // os.Chdir(repoDir)
  // if err != nil {
  //   fmt.Printf("Failed to execute command '%s'. %s %s\n", command, output, err.Error())
  //   os.Exit(1)
  // }

  // // Exec 'yarn run build'
  // command = fmt.Sprintf("yarn run bundle")
  // fmt.Printf("command: '%s'\n", command)
  // os.Chdir(webDir)
  // output, err = exec.Command("bash", "-c", command).Output()
  // os.Chdir(repoDir)
  // if err != nil {
  //   fmt.Printf("Failed to execute command '%s'. %s\n", command, output, err.Error())
  //   os.Exit(1)
  // }

  // // Create the tarball directory if it does not exist
  // _, err = os.Stat(tarballDir)
  // if os.IsNotExist(err) {
  //   os.MkdirAll(tarballDir, 0700)
  // }

  // // If the tarball exists, remove it
  // _, err = os.Stat(tarballFilepath)
  // if !os.IsNotExist(err) {
  //   os.RemoveAll(tarballFilepath)
  //   if err != nil {
  //     fmt.Printf("Error removing tarball. %s\n", err.Error())
  //     os.Exit(1)
  //   }
  // }

  // // Tar the pgadmin directory
  // parentDir := filepath.Dir(repoDir)
  // command = fmt.Sprintf("tar -C  %s -czf %s toolbox-pgadmin4", parentDir, tarballFilepath)
  // fmt.Printf("command: '%s'\n", command)
  // _, err = exec.Command("bash", "-c", command).Output()
  // if err != nil {
  //   fmt.Printf("error tarring pgadmin. %s\n", err.Error())
  //   os.Exit(1)
  // }
}

func main() {
  var builder Builder
  builder.build()
}
