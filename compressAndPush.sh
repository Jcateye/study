#!/usr/bin/env bash
targetPath=''
unit=$(expr 1024 \* 1024)
gitDir=''
fileName='filename.tar.gz'
while getopts ":a:b:c:d" opt; do
  # shellcheck disable=SC2220
  case $opt in
  a)
    echo "压缩文件: $OPTARG"
    targetPath=$OPTARG
    ;;
  b)
    echo "每个文件大小:$OPTARG""B"
    unit=$(expr $OPTARG)
    ;;
  c)
    echo "git目录: $OPTARG"
    gitDir=$OPTARG
    curPath=$(pwd)
    cd $gitDir
    gitDir=$(pwd)
    cd $curPath
    ;;
  d)
    echo "参数d:$OPTARG"
    fileName=$OPTARG
    ;;
  esac
done
if test -z $gitDir; then
  echo "-c git目录必须指定"
  exit 1
fi
pwd
tar -zcvf 'filename.tar.gz' $targetPath
split -b ${unit} -d -a 1 'filename.tar.gz' 'filename.tar.gz.'
curPath=$(pwd)
rm -f 'filename.tar.gz'
mv "filename.tar.gz."* $gitDir
cd $gitDir
pwd
listFileStr=$(ls $fileName*)
echo $listFileStr
for file in $listFileStr; do
  git add $file
  git status
  git commit -m 'add'$file
  git push origin master
done
