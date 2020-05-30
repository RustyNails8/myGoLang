
<!-- @import "[TOC]" {cmd="toc" depthFrom=1 depthTo=6 orderedList=false} -->
# RustyNailsGoNotes

## Build Strategy
### Head First Go Book samples

* Create module
``` go mod init gitlab.com/RustyNails8/<pkgFolder>/<pkg.go> ```

* While on cmd folder, run
``` gBUILD.cmd pakgFolder\pakgModule\pakgmain.go ```

* Verb Output
``` 
%f Floating-point number
%d Decimal integer
%s String
%t Boolean (true or false)
%v Any value (chooses an appropriate format based on the supplied value’s type)
%#v Any value, formatted as it would appear in Go program code
%T Type of the supplied value (int, string, etc.)
%% A literal percent sign 
```