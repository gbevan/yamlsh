#!./yamlsh --yaml=${MYFILE}
echo "IN SCRIPT"

env | grep "$YAMLSH_PREFIX"

exit 0
