if [ $# -ne 1 ]; then
	echo "USAGE: $0 <action>"
	echo " action: build, build-native, build-all, vendor"
fi

case "${1}" in
	build) go build;;
	build-native) cd lib; bazel build testcpplib; cp -f bazel-bin/testcpplib/libtestcpplib.a ../bin/; cd -;;
	build-all) $0 build-native; $0 build;;
	vendor) go vendor;;
esac
