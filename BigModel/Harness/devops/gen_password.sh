export DEEPSEEK_HERNESS_AUTH="sxk123456"
docker run --rm httpd:2.4-alpine htpasswd -nb admin ${DEEPSEEK_HERNESS_AUTH} > ./config/htpasswd
echo "./config/htpasswd" >> .gitignore
echo ".env" >> .gitignore