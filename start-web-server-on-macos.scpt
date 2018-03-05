# Copyright 2018 <chaishushan{AT}gmail.com>. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

tell application "Finder"
	set currentDir to POSIX path of ((container of (path to me)) as text)
end tell

tell application "Terminal" to activate

tell application "Terminal"
	do script ("cd '" & currentDir & "'")
	do script ("python -m SimpleHTTPServer 8081") in front window
end tell
