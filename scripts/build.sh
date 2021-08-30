go build -ldflags "-s -w" -o bin/passgen cmd/*.go &

pid=$!

directory="bin"

if [[ -d $directory ]]; then
  while ps -p $pid &>/dev/null ; do
      echo -ne 'Building... [****              ](25%)\r'
      sleep .3
      echo -ne 'Building... [********          ](50%)\r'
      sleep .3
      echo -ne 'Building... [***********       ](75%)\r'
      sleep .3
      echo -ne 'Building... [***************** ](99%)\r'
      sleep .3
  done
  echo -ne 'Building... [******************](100%)\r'
else
  mkdir bin
  while ps -p $pid &>/dev/null ; do
      echo -ne 'Building... [****              ](25%)\r'
      sleep .3
      echo -ne 'Building... [********          ](50%)\r'
      sleep .3
      echo -ne 'Building... [***********       ](75%)\r'
      sleep .3
      echo -ne 'Building... [***************** ](99%)\r'
      sleep .3
  done
  echo -ne 'Building... [******************](100%)\r'
fi

