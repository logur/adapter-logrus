#!/bin/bash

# Destination directory of modifications
DEST="."

# Original project variables
originalProjectName="adapter-template"
originalAdapterName="template"

# Prepare testing
if [[ ! -z "${TEST}" ]]; then
    #set -xe
    DEST="tmp/inittest"
    mkdir -p ${DEST}
    echo "." > tmp/.gitignore
fi

function prompt() {
    echo -n -e "\033[1;32m?\033[0m \033[1m$1\033[0m ($2) "
}

function replace() {
    if [[ ! -z "${TEST}" ]]; then
        dest=$(echo $2 | sed "s|^${DEST}/||")
        mkdir -p $(dirname "${DEST}/${dest}")
        if [[ "$2" == "${DEST}/${dest}" ]]; then
            sed -E -e "$1" $2 > ${DEST}/${dest}.new
            mv -f ${DEST}/${dest}.new ${DEST}/${dest}
        else
            sed -E -e "$1" $2 > ${DEST}/${dest}
        fi
    else
        sed -E -e "$1" $2 > $2.new
        mv -f $2.new $2
    fi
}

function move() {
    if [[ ! -z "${TEST}" ]]; then
        dest=$(echo $2 | sed "s|^${DEST}/||")
        mkdir -p $(dirname "${DEST}/${dest}")
        cp -r "$1" ${DEST}/${dest}
    else
        mv $@
    fi
}

function remove() {
    if [[ -z "${TEST}" ]]; then
        rm $@
    fi
}

defaultAdapterName="template"
prompt "Adapter name" ${defaultAdapterName}
read adapterName
adapterName=$(echo "${adapterName:-${defaultAdapterName}}" | sed 's/[[:space:]]//g')

prompt "Update README" "Y/n"
read updateReadme
updateReadme=${updateReadme:-y}

prompt "Remove init script" "y/N"
read removeInit
removeInit=${removeInit:-n}

# IDE configuration
move .idea/${originalProjectName}.iml .idea/adapter-${adapterName}.iml
replace "s|.idea/${originalProjectName}.iml|.idea/adapter-${adapterName}.iml|g" .idea/modules.xml

# Run configurations
replace 's|name="'${originalProjectName}'"|name="'adapter-${adapterName}'"|' .idea/runConfigurations/Test.xml

# Other project files
declare -a files=("CHANGELOG.md" "go.mod" ".golangci.yml")
for file in "${files[@]}"; do
    if [[ -f "${file}" ]]; then
        replace "s|${originalAdapterName}|${adapterName}|" ${file}
    fi
done

if [[ "${removeInit}" != "n" && "${removeInit}" != "N" ]]; then
    remove "$0"
fi

# Update readme
if [[ "${updateReadme}" == "y" || "${updateReadme}" == "Y" ]]; then
    replace "s|${originalProjectName}|adapter-${adapterName}|g; s|/${originalAdapterName}|/${adapterName}|g; s|${originalAdapterName}adapter|${adapterName}adapter|g" README.md
fi
