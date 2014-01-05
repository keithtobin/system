package main

import (
//   "os/exec"
   "fmt"
)

type IDeploymentEngine interface {
        InstallPackage(packageName string) (err error)
        InstallPackages(packageNames []string) (err error)
        AddRepository(repositoryName string, triggerUpdate bool) (err error)
        AddRepositorys(repositoryNames []string, triggerUpdate bool) (err error)
}

type DeploymentEngine struct {}

func New() (ide IDeploymentEngine) {
        de := DeploymentEngine{}
        fed := IDeploymentEngine(de)
        return fed
}


func (DeploymentEngine) InstallPackage(packageName string) (err error) {

        s := fmt.Sprintf("DEBIAN_FRONTEND=noninteractive apt-get install -y %s", packageName)

        return runCommand(s)
}

func (DeploymentEngine) InstallPackages(packageNames []string) (err error) {

        for i := range packageNames {
                s := fmt.Sprintf("DEBIAN_FRONTEND=noninteractive apt-get install -y %s", packageNames[i])
                if err := runCommand(s); err != nil {
                        return err
                }
        }

        return
}



func (DeploymentEngine) AddRepository(repositoryName string, triggerUpdate bool) (err error) {

        s := fmt.Sprintf("DEBIAN_FRONTEND=noninteractive add-apt-repository %s", repositoryName)
        if err := runCommand(s); err != nil {
                return err
        }

        if triggerUpdate == true {

                if err := runCommand("DEBIAN_FRONTEND=noninteractive apt-get update"); err != nil {
                        return err
                }
        }

        return

}

func (DeploymentEngine) AddRepositorys(repositoryNames []string, triggerUpdate bool) (err error) {

        for i := range repositoryNames {
                s := fmt.Sprintf("DEBIAN_FRONTEND=noninteractive add-apt-repository %s", repositoryNames[i])
                if err := runCommand(s); err != nil {
                        return err
                }
        }

        if triggerUpdate == true {

                if err := runCommand("DEBIAN_FRONTEND=noninteractive apt-get update"); err != nil {
                        return err
                }
        }

        return

}

func runCommand(cmd string) (err error) {

        //out, err := exec.Command(cmd).CombinedOutput()
        //if err != nil {
        //      return err
        //}

        fmt.Printf("RunningCommand: %s\n",cmd)
        //fmt.Printf("%s",out)
        return
}

func (DeploymentEngine) ConfigureNTP(url string) (err error) {


        return
}

func (DeploymentEngine) ConfigureHostname(name string) (err error){


        return
}

func (DeploymentEngine) RestartNetworking() (err error){

        return
}

func (DeploymentEngine) Reboot() (err error){

        return
}



