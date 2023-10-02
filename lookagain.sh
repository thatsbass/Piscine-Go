find . -type f -name "*.sh" -execdir basename {} \; | sort -n -r| sed -e 's/\.sh$//'
