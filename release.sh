git add .
git commit -m 'finish mini mq'
git push origin develop
git checkout master
git pull
git merge develop
git push origin master
git tag v1.0.0
git push origin v1.0.0