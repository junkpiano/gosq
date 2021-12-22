package gosq

//go:generate gojson -o component.go -name "ComponentResponse" -pkg "gosq" -input json/component.json
//go:generate gojson -o system_info.go -name "SystemInfo" -pkg "gosq" -input json/systeminfo.json
//go:generate gojson -o branch_list.go -name "BranchList" -pkg "gosq" -input json/branch_list.json
//go:generate gojson -o projects.go -name "Projects" -pkg "gosq" -input json/projects.json
//go:generate gojson -o alm_settings.go -name "AlmSettings" -pkg "gosq" -input json/alm_settings.json
