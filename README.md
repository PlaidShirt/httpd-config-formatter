# httpd-config-formatter
A simple tool to format Apache httpd config files at the command line.

**Installation**
Assuming you have a working go installion type: go get github.com/plaidshirt/httpd-config-formatter

**Usage**
Typing "httpd-config-formatter -h" will give usage information.  With no arguments the command will look for a file called httpd.conf in the current directory and pipe a formatted version of that file to standard output.

The following flags are available:

-h Print usage information.

-i Format the file "in place".  A backup copy of the file with a ".bak" extension will be made and the existing file will be formatted.

-f Specify the file to process (defaults to httpd.conf).

-p Specifies the number of pad characters to use (default 4).

**Examples**

The following will back up the httpd.conf file in the current directory to httpd.conf.bak and then create a formatted version of the orginal file:

&nbsp;&nbsp;&nbsp;httpd-config-formatter -i

This example will back up the ssl.conf file to ssl.conf.bak and then create a formatted version in ssl.conf, using two spaces per level of indentation:

&nbsp;&nbsp;&nbsp;httpd-config-formatter -i -p 2 -f ssl.conf

This final example will take the file passed in on the command line and print the formatted output to stdout (ssl.conf in this case):

&nbsp;&nbsp;&nbsp;httpd-config-formatter ssl.conf

**Future Enhancements**
Add a flag to select tab vs space for the padding character.
Add a flat to change the backup file extension

**Disclaimer**
No warranty of any kind is expressed or implied with this software.  Always back up your data and be sure to have a recovery plan.  If by using this software you accidentally destroy your system all I can offer is a heartfelt apology.