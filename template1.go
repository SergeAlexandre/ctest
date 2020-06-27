package main


var template1 = `
#!/usr/bin/env bash

echo "=========================== {{.Name}}"

set +x

failed=()

appendFailed() {
  msg="$1/$2: $3"
  echo "FAILED: $msg"
  failed+=("$msg")
{{- if .ExitOnFail }}
  exit 1
{{- end }}
}


setup() {
    :
{{- range .Setup}}
    {{.}}
{{- end}}
}

tearDown() {
    :
{{- range .TearDown}}
    {{.}}
{{- end}}
}

set -e
# Global setup
{{- range .GlobalSetup}}
{{.}}
{{- end}}

set +e
{{- range .Tests}}
echo "--------------------------- {{.Name}}"
TESTNAME="{{.Name}}"
setup
{{- range .Steps}}

echo "---------- {{.Name}}"
STEPNAME="{{.Name}}"
output=$({{.Command}}); retCode=$?
{{- if .Echo }}
echo "----- $output"
{{- end}}
{{- if .RetCode }}
if [ $retCode -ne {{.RetCode}} ]; then appendFailed "$TESTNAME" "$STEPNAME" "retCode is '$retCode' Should be '{{.RetCode}}'"; fi
{{- end}}
{{- if .Output }}
if [ "$output" != "{{.Output}}" ]; then appendFailed "$TESTNAME" "$STEPNAME" "output is '$output' Should be '{{.Output}}'"; fi
{{- end}}
{{- if .OutputExp }}
if [[ ! "$output" =~ {{.OutputExp}} ]]; then appendFailed "$TESTNAME" "$STEPNAME" "output is '$output' Should match '{{.OutputExp}}'"; fi
{{- end}}
{{- end}}
tearDown
{{- end}}



# Global tearDown
{{- range .GlobalTearDown}}
{{.}}
{{- end}}

echo ""
echo "${#failed[@]} failure(s) on {{ .Name }}:"

printf '%s\n' "${failed[@]}"

exit ${#failed[@]}

`
