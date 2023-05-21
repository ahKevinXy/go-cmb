


time1=$(date "+%Y-%m-%d %H:%M:%S")
time2="**$1**  *${time1}*"

name=$(git config user.name)
commit="${time2}"


content="\n  $name ${time2} "
content1="NR==2{print \"$content\"}1"

echo $content1
if test $2 != "n"
then
awk "$content1" ./docs/CHANGELOG.md >>./docs/CHANGELOG_new.md
#awk "$content1" ./CHANGELOG.md >>./CHANGELOG.md

mv ./docs/CHANGELOG_new.md ./docs/CHANGELOG.md
#mv ./README_new.md ./README.md



fi
# 提交代码
git add .
git commit -m"$1"
git push