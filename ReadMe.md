## Database Script
Using Visual Studio 2017 Sql Server Data Tools (SSDT) to do a schema compare of the FGB_Portal to FGB_Portal_Trunk.

* **BACKUP FGBPortalTrunk on 10.146.200.205(FGBDEV05)**
* Be sure you are comparing from 'FGB_Portal to FGB_Portal_Trunk on 10.146.200.205(FGBDEV05)
* In options, check "Generate Smart Defaults When Applicable"
* Run Compare (if it fails to intialise, you may need to re-select the sources)  (If you encounter “Target is unavailable” Error, under the source connection, check advanced settings > the authentication is Sqlpasasword and network library is TCP/IP)
* Normally keep all the Delete items checked. Else we can get naming conflicts on constraints, etc.
* Generate Script. 
* Open the Update Template in $/Portal/Releases/Update Scripts/Update Script Template.sql, and copy the script in the script block.
* Update the Version Number ie: `:setvar VersionNumber = "2.2"`
* Test script against FGB_Portal_Trunk to verify the script works and make sure FGB_Portal_Trunk is up to date.
* Save script as {{Year}}\{VersionNumber}}\UpgradeScript.sql ie:  $/Portal/Releases/Update Scripts/2018_2.2_UpgradeScript.sql. 
* Save file to $/Portal/Releases/Update Scripts
* Add File to Source Control and check in.

## Merge
For each of the Projects (External, Services and Portal), merge the code from Development branch to Trunk branch and check in.

Inside VS 2017, open up Source Control Explorer.

* Open the Project folder, and right click on Development branch > Branching and Merging.. > Merge
* Resolve any conflicts **MAKE SURE TO NOT MERGE TO DEVELOPMENT!**
* In pending check ins, we need all the ids of the resolved tasks and bugs. See the next step. Ignore this for External.
* Inside Team Services > Portal > Work > Queries > On 'Output Work item Ids', select More Actions(the ...), click Enhanced Query Export. Select all the ids, and edit it to have them all comma seperated and in a single line. Add these to Related Work Items.
* When checking in make certain that only files for that Project, for Trunk, are being checked in. **Don't check in Development files!**
* Check in the files.

## Release
For Services and Portal a manual release is required.

* Open up Team Services
* Project > Build and release > Library > Version Control Variables > Increment Minor Version
* Builds > Build Definitions > Trunk > Trunk - Services > Queue New Build
* Once completed, download the artifacts (webframework in Services. Admin, Content, Portal, Portal2 in Portal)

## Mainshare
Copy the artifacts and update script to Portal Builds