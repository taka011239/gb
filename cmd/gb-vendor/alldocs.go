// DO NOT EDIT THIS FILE.
//go:generate gb vendor help documentation

/*
gb-vendor, a gb plugin to manage your vendored dependencies.

Usage:

        gb vendor command [arguments]

The commands are:

        fetch       fetch a remote dependency
        update      update a local dependency
        list        lists dependencies, one per line
        delete      deletes a local dependency
        purge       purges all unreferenced dependencies

Use "gb vendor help [command]" for more information about a command.

Additional help topics:


Use "gb vendor help [topic]" for more information about that topic.


Fetch a remote dependency

Usage:

        gb vendor fetch [-branch branch | -revision rev | -tag tag | -precaire] importpath

fetch vendors the upstream import path.

Flags:
	-branch branch
		fetch from the name branch. If not supplied the default upstream
		branch will be used.
	-tag tag
		fetch the specified tag. If not supplie the default upstream
		branch will be used.
	-revision rev
		fetch the specific revision from the branch (if supplied). If no
		revision supplied, the latest available will be supplied.
	-precaire
		allow the use of insecure protocols.


Update a local dependency

Usage:

        gb vendor update [-all] import

gb vendor update will replaces the source with the latest available from the head of the master branch.

Updating from one copy of a dependency to another comes with several restrictions.
The first is you can only update to the head of the branch your dependency was vendered from, switching branches is not supported.
The second restriction is if you have used -tag or -revision while vendoring a dependency, your dependency is "headless"
(to borrow a term from git) and cannot be updated.

To update across branches, or from one tag/revision to another, you must first use gb vendor delete to remove the dependency, then
gb vendor fetch [-tag | -revision | -branch] to replace it.

Flags:
	-all
		will update all dependencies in the manifest, otherwise only the dependency supplied.
	-precaire
		allow the use of insecure protocols.


Lists dependencies, one per line

Usage:

        gb vendor list [-f format]

gb vendor list formats lists the contents of the manifest file.

The output

Flags:
	-f
		controls the template used for printing each manifest entry. If not supplied
		the default value is "{{.Importpath}}\t{{.Repository}}{{.Path}}\t{{.Branch}}\t{{.Revision}}"


Deletes a local dependency

Usage:

        gb vendor delete [-all] importpath

delete removes a dependency from $PROJECT/vendor/src and the vendor manifest

Flags:
	-all
		remove all dependencies


Purges all unreferenced dependencies

Usage:

        gb vendor purge

gb vendor purge will remove all unreferenced dependencies


*/
package main
