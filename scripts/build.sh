#!bin/bash

start=$(date +%s.%N)

go build -ldflags "-s -w" -o bin/passgen cmd/*.go &

pid=$!

directory="bin"

green_color='\e[32m'
light_grey_color='\e[37m'


BAR="█"

bar_loaded_length=3
bar_unloaded_length=$(($bar_loaded_length * 10))

inc_bar_item="$green_color█"
left_bar_item="$light_grey_color█"

total_size=0
size_value_counter=0

progress_bar() {
	percent_bar=$(("$1*100/${total_size}*100"/100))

	loaded_bar=$(("${percent_bar}*${bar_loaded_length}"/10))
 
	unloaded_bar=$(($bar_unloaded_length-$loaded_bar))
	
	completed=$(printf "%${loaded_bar}s")

	left_bar=$(printf "%${unloaded_bar}s")

	echo -ne "Building... ${completed// /$inc_bar_item}${BAR:i++%${#BAR}:1}${left_bar// /$left_bar_item} (${percent_bar}%)\r"
}

if [[ -d $directory ]]; then
	while ps -p $pid &>/dev/null ; do
		total_size=$(($total_size+1))
	done

	while [[ ${size_value_counter} -lt ${total_size} ]] ; do
		sleep .1
		size_value_counter=$(($size_value_counter+1))
		if [[ "${size_value_counter}" == "${total_size}" ]]; then
			BAR="*"
		fi

		progress_bar $size_value_counter
	done
else
	mkdir bin
	while ps -p $pid &>/dev/null ; do
                total_size=$(($total_size+1))
        done

        while [[ ${size_value_counter} -lt ${total_size} ]] ; do
                 sleep .1
                 size_value_counter=$(($size_value_counter+1))
                 if [[ "${size_value_counter}" == "${total_size}" ]]; then
                      BAR="*"
                fi

                 progress_bar $size_value_counter
         done 
fi
