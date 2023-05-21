


time1=$(date "+%Y-%m-%d %H:%M:%S")
time2=" *${time1}*  **$1** "

name=$(git config user.name)
commit="${time2}"


content="\n ${time2}   $name  "
content1="NR==2{print \"$content\"}1"



echo $2
if test $2 != "n"
then
awk "$content1" ./docs/CHANGELOG.md >>./docs/CHANGELOG_new.md
#awk "$content1" ./CHANGELOG.md >>./CHANGELOG.md

mv ./docs/CHANGELOG_new.md ./docs/CHANGELOG.md
#mv ./README_new.md ./README.md

cp -rf  ./docs/CHANGELOG.md ./CHANGELOG.md

fi
# 提交代码
git add .
git commit -m"${time1}  $1"
git push