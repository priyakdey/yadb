#!/usr/bin/env python3

# Copyright (c) 2023 Priyak Dey
#
# Permission is hereby granted, free of charge, to any person obtaining a copy
# of this software and associated documentation files (the “Software”),
# to deal in the Software without restriction, including without limitation the rights
# to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
# the Software, and to permit persons to whom the Software is furnished to do so,
# subject to the following conditions:
#
# The above copyright notice and this permission notice shall be included in all copies
# or substantial portions of the Software.
#
# THE SOFTWARE IS PROVIDED “AS IS”, WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED,
# INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR
# PURPOSE AND NONINFRINGEMENT.
# IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM,
# DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE,
# ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS
# IN THE SOFTWARE.


###########################################################################################
#####                                                                                 #####
#####           Helper script to create new source files with LICENSE info            #####
#####                                                                                 #####
###########################################################################################


# -*- coding: utf-8 -*-

from datetime import datetime
import os
import sys


license_tmpl ="""// Copyright (c) {{ Year }} Priyak Dey
// 
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the “Software”),
// to deal in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
// the Software, and to permit persons to whom the Software is furnished to do so,
// subject to the following conditions:
// 
// The above copyright notice and this permission notice shall be included in all copies
// or substantial portions of the Software.
// 
// THE SOFTWARE IS PROVIDED “AS IS”, WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED,
// INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR
// PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM,
// DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE,
// ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS
// IN THE SOFTWARE.
"""

file_content_tmpl = """{{ License }}
package {{ PackageName }}
"""

if (len(sys.argv) < 3):
	print("ERROR: missing mandatory arguments.")
	print("Usage hints:")
	print("\t./cf.py <package_name> <file_path>")
	print("Examples:")
	print("\t./cf.py main main.go", "\t\t\tCreates a file main.go with package main in ./ folder")
	print("\t./cf.py util internal/util/httputil.go", "\tCreates a file httputil.go with package util in ./internal/util/ folder")
	exit(1)

package_name_arg = sys.argv[1]
file_path_arg = sys.argv[2]

curr_year = str(datetime.now().year)
license = license_tmpl.replace("{{ Year }}", curr_year)
file_content = file_content_tmpl.replace("{{ License }}", license).replace("{{ PackageName }}", package_name_arg)

*file_path_list, file_name = file_path_arg.split("/")

file_path = os.path.join(os.path.dirname(__file__), *file_path_list)

try:
	os.makedirs(file_path)
except OSError:
	# ignore if directory exists. We can create multiple files inside same dir
	pass	

file = os.path.join(file_path, file_name)

if os.path.exists(file):
	print(f"ERROR: {file} does exist.")
	exit(1)

with open(os.path.join(file_path, file_name), mode="w") as fp:
	fp.write(file_content)
