---
layout: post
title: "vscode launch.json"
subtitle: ''
author: "Terry"
header-style: text
tags:
  - vscode
---

`${workspaceFolder}` :表示当前workspace文件夹路径，

`${workspaceRootFolderName}`:表示workspace的文件夹名，

`${file`}:文件自身的绝对路径，

`${relativeFile}`:文件在workspace中的路径

`${fileBasenameNoExtension}`:当前文件的文件名，不带后缀

`${fileBasename}`:当前文件的文件名

`${fileDirname}`:文件所在的文件夹路径

`${fileExtname}`:当前文件的后缀

`${lineNumber}`:当前文件光标所在的行号

`${env:PATH}`:系统中的环境变量