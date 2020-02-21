package pkgmanlib

// Update this function is for updating all packages and services. It will check if the user is a user or if it is root
// (if root, no sudo, if !root then add sudo), before it will check the oslib for the correct OS, to create the correct command
func Update(username string, os string) (updatecommand string) {
	if username == "root" {
		updatecommand = ""
	} else {
		updatecommand = "sudo"
	}
	updatecommand = updatecommand + PkgRefresh[os] + ";" + PkgUpdate[os] + "2>&1"
	return updatecommand
}

// UpdateOS same as the Update function, but will run any OS/distribution related upgrade/update parameters to upgrade all packages including the OS
// some OS's does not have this functionality and updates the OS during a normal update. So if the OS does have a separate upgrade option
// (eg. debian dist-upgrade), then this function is redundant
func UpdateOS(username string, os string) (updatecommand string) {
	if username == "root" {
		updatecommand = ""
	} else {
		updatecommand = "sudo"
	}
	updatecommand = updatecommand + PkgRefresh[os] + ";" + PkgUpdateOS[os] + "2>&1"
	return updatecommand
}

// Install this will install any packages specified on the servers, creating the correct command for each major distribution and/or package manager
func Install(username string, os string, cmdargs string) (installcommand string) {
	if username == "root" {
		installcommand = ""
	} else {
		installcommand = "sudo"
	}
	installcommand = installcommand + PkgRefresh[os] + ";" + PkgInstall[os] + cmdargs + " -y 2>&1"
	return installcommand
}

// Uninstall this will install any packages specified on the servers, creating the correct command for each major distribution and/or package manager
func Uninstall(username string, os string, cmdargs string) (uninstallcommand string) {
	if username == "root" {
		uninstallcommand = ""
	} else {
		uninstallcommand = "sudo"
	}
	uninstallcommand = uninstallcommand + PkgRefresh[os] + ";" + PkgUninstall[os] + cmdargs + " -y 2>&1"
	return uninstallcommand
}
