package gosq

type SystemInfo struct {
	Database struct {
		Database                               string `json:"Database"`
		Database_Version                       string `json:"Database Version"`
		Driver                                 string `json:"Driver"`
		Driver_Version                         string `json:"Driver Version"`
		Pool_Active_Connections                int64  `json:"Pool Active Connections"`
		Pool_Idle_Connections                  int64  `json:"Pool Idle Connections"`
		Pool_Initial_Size                      int64  `json:"Pool Initial Size"`
		Pool_Max_Connections                   int64  `json:"Pool Max Connections"`
		Pool_Max_Idle_Connections              int64  `json:"Pool Max Idle Connections"`
		Pool_Max_Wait__ms                      int64  `json:"Pool Max Wait (ms)"`
		Pool_Min_Idle_Connections              int64  `json:"Pool Min Idle Connections"`
		Pool_Remove_Abandoned                  bool   `json:"Pool Remove Abandoned"`
		Pool_Remove_Abandoned_Timeout__seconds int64  `json:"Pool Remove Abandoned Timeout (seconds)"`
		URL                                    string `json:"URL"`
		Username                               string `json:"Username"`
		Version_Status                         string `json:"Version Status"`
	} `json:"Database"`
	ElasticSearch struct {
		Indices struct {
			Issues struct {
				Docs       int64  `json:"Docs"`
				Shards     int64  `json:"Shards"`
				Store_Size string `json:"Store Size"`
			} `json:"issues"`
			Logs struct {
				Docs       int64  `json:"Docs"`
				Shards     int64  `json:"Shards"`
				Store_Size string `json:"Store Size"`
			} `json:"logs"`
			Rules struct {
				Docs       int64  `json:"Docs"`
				Shards     int64  `json:"Shards"`
				Store_Size string `json:"Store Size"`
			} `json:"rules"`
			Sourcelines struct {
				Docs       int64  `json:"Docs"`
				Shards     int64  `json:"Shards"`
				Store_Size string `json:"Store Size"`
			} `json:"sourcelines"`
			Users struct {
				Docs       int64  `json:"Docs"`
				Shards     int64  `json:"Shards"`
				Store_Size string `json:"Store Size"`
			} `json:"users"`
			Views struct {
				Docs       int64  `json:"Docs"`
				Shards     int64  `json:"Shards"`
				Store_Size string `json:"Store Size"`
			} `json:"views"`
		} `json:"Indices"`
		Nodes struct {
			SixDYkdNuvT6abEdShyEK5ow struct {
				Address             string `json:"Address"`
				Disk_Available      string `json:"Disk Available"`
				Disk_Usage          string `json:"Disk Usage"`
				Field_Cache_Memory  string `json:"Field Cache Memory"`
				Filter_Cache_Memory string `json:"Filter Cache Memory"`
				ID_Cache_Memory     string `json:"ID Cache Memory"`
				JVM_Heap_Max        string `json:"JVM Heap Max"`
				JVM_Heap_Usage      string `json:"JVM Heap Usage"`
				JVM_Heap_Used       string `json:"JVM Heap Used"`
				JVM_Non_Heap_Used   string `json:"JVM Non Heap Used"`
				JVM_Threads         int64  `json:"JVM Threads"`
				Open_Files          int64  `json:"Open Files"`
				Query_Cache_Memory  string `json:"Query Cache Memory"`
				Store_Size          string `json:"Store Size"`
				Type                string `json:"Type"`
			} `json:"6DYkdNuvT6abEdShyEK5ow"`
		} `json:"Nodes"`
		Number_of_Nodes int64  `json:"Number of Nodes"`
		State           string `json:"State"`
	} `json:"ElasticSearch"`
	Health        string        `json:"Health"`
	Health_Causes []interface{} `json:"Health Causes"`
	JvmProperties struct {
		Awt_toolkit                                           string `json:"awt.toolkit"`
		Catalina_base                                         string `json:"catalina.base"`
		Catalina_home                                         string `json:"catalina.home"`
		Catalina_useNaming                                    string `json:"catalina.useNaming"`
		File_encoding                                         string `json:"file.encoding"`
		File_encoding_pkg                                     string `json:"file.encoding.pkg"`
		File_separator                                        string `json:"file.separator"`
		Ftp_nonProxyHosts                                     string `json:"ftp.nonProxyHosts"`
		GopherProxySet                                        string `json:"gopherProxySet"`
		HTTP_agent                                            string `json:"http.agent"`
		HTTP_nonProxyHosts                                    string `json:"http.nonProxyHosts"`
		Java_awt_graphicsenv                                  string `json:"java.awt.graphicsenv"`
		Java_awt_headless                                     string `json:"java.awt.headless"`
		Java_awt_printerjob                                   string `json:"java.awt.printerjob"`
		Java_class_path                                       string `json:"java.class.path"`
		Java_class_version                                    string `json:"java.class.version"`
		Java_endorsed_dirs                                    string `json:"java.endorsed.dirs"`
		Java_ext_dirs                                         string `json:"java.ext.dirs"`
		Java_home                                             string `json:"java.home"`
		Java_io_tmpdir                                        string `json:"java.io.tmpdir"`
		Java_library_path                                     string `json:"java.library.path"`
		Java_net_preferIPv4Stack                              string `json:"java.net.preferIPv4Stack"`
		Java_runtime_name                                     string `json:"java.runtime.name"`
		Java_runtime_version                                  string `json:"java.runtime.version"`
		Java_specification_name                               string `json:"java.specification.name"`
		Java_specification_vendor                             string `json:"java.specification.vendor"`
		Java_specification_version                            string `json:"java.specification.version"`
		Java_vendor                                           string `json:"java.vendor"`
		Java_vendor_url                                       string `json:"java.vendor.url"`
		Java_vendor_url_bug                                   string `json:"java.vendor.url.bug"`
		Java_version                                          string `json:"java.version"`
		Java_vm_info                                          string `json:"java.vm.info"`
		Java_vm_name                                          string `json:"java.vm.name"`
		Java_vm_specification_name                            string `json:"java.vm.specification.name"`
		Java_vm_specification_vendor                          string `json:"java.vm.specification.vendor"`
		Java_vm_specification_version                         string `json:"java.vm.specification.version"`
		Java_vm_vendor                                        string `json:"java.vm.vendor"`
		Java_vm_version                                       string `json:"java.vm.version"`
		Line_separator                                        string `json:"line.separator"`
		Org_apache_catalina_startup_EXITONINITFAILURE         string `json:"org.apache.catalina.startup.EXIT_ON_INIT_FAILURE"`
		Org_apache_tomcat_util_buf_UDecoder_ALLOWENCODEDSLASH string `json:"org.apache.tomcat.util.buf.UDecoder.ALLOW_ENCODED_SLASH"`
		Os_arch                                               string `json:"os.arch"`
		Os_name                                               string `json:"os.name"`
		Os_version                                            string `json:"os.version"`
		Path_separator                                        string `json:"path.separator"`
		SocksNonProxyHosts                                    string `json:"socksNonProxyHosts"`
		Sun_arch_data_model                                   string `json:"sun.arch.data.model"`
		Sun_boot_class_path                                   string `json:"sun.boot.class.path"`
		Sun_boot_library_path                                 string `json:"sun.boot.library.path"`
		Sun_cpu_endian                                        string `json:"sun.cpu.endian"`
		Sun_cpu_isalist                                       string `json:"sun.cpu.isalist"`
		Sun_io_unicode_encoding                               string `json:"sun.io.unicode.encoding"`
		Sun_java_command                                      string `json:"sun.java.command"`
		Sun_java_launcher                                     string `json:"sun.java.launcher"`
		Sun_jnu_encoding                                      string `json:"sun.jnu.encoding"`
		Sun_management_compiler                               string `json:"sun.management.compiler"`
		Sun_nio_ch_bugLevel                                   string `json:"sun.nio.ch.bugLevel"`
		Sun_os_patch_level                                    string `json:"sun.os.patch.level"`
		User_country                                          string `json:"user.country"`
		User_country_format                                   string `json:"user.country.format"`
		User_dir                                              string `json:"user.dir"`
		User_home                                             string `json:"user.home"`
		User_language                                         string `json:"user.language"`
		User_name                                             string `json:"user.name"`
		User_timezone                                         string `json:"user.timezone"`
	} `json:"JvmProperties"`
	Plugins struct {
		Findbugs struct {
			Name    string `json:"Name"`
			Version string `json:"Version"`
		} `json:"findbugs"`
		Java struct {
			Name    string `json:"Name"`
			Version string `json:"Version"`
		} `json:"java"`
		Scmgit struct {
			Name    string `json:"Name"`
			Version string `json:"Version"`
		} `json:"scmgit"`
		Scmsvn struct {
			Name    string `json:"Name"`
			Version string `json:"Version"`
		} `json:"scmsvn"`
	} `json:"Plugins"`
	Search_State struct {
		CPU_Usage                             int64  `json:"CPU Usage (%)"`
		Disk_Available                        string `json:"Disk Available"`
		Field_Data_Circuit_Breaker_Estimation string `json:"Field Data Circuit Breaker Estimation"`
		Field_Data_Circuit_Breaker_Limit      string `json:"Field Data Circuit Breaker Limit"`
		Field_Data_Memory                     string `json:"Field Data Memory"`
		JVM_Heap_Max                          string `json:"JVM Heap Max"`
		JVM_Heap_Usage                        string `json:"JVM Heap Usage"`
		JVM_Heap_Used                         string `json:"JVM Heap Used"`
		JVM_Non_Heap_Used                     string `json:"JVM Non Heap Used"`
		JVM_Threads                           int64  `json:"JVM Threads"`
		Max_File_Descriptors                  int64  `json:"Max File Descriptors"`
		Open_File_Descriptors                 int64  `json:"Open File Descriptors"`
		Query_Cache_Memory                    string `json:"Query Cache Memory"`
		Request_Cache_Memory                  string `json:"Request Cache Memory"`
		Request_Circuit_Breaker_Estimation    string `json:"Request Circuit Breaker Estimation"`
		Request_Circuit_Breaker_Limit         string `json:"Request Circuit Breaker Limit"`
		State                                 string `json:"State"`
		Store_Size                            string `json:"Store Size"`
		Translog_Size                         string `json:"Translog Size"`
	} `json:"Search State"`
	SonarQube struct {
		Allow_Users_to_Sign_Up       bool   `json:"Allow Users to Sign Up"`
		Automatic_User_Creation      bool   `json:"Automatic User Creation"`
		Data_Dir                     string `json:"Data Dir"`
		External_User_Authentication string `json:"External User Authentication"`
		Force_authentication         bool   `json:"Force authentication"`
		Home_Dir                     string `json:"Home Dir"`
		Logs_Dir                     string `json:"Logs Dir"`
		Temp_Dir                     string `json:"Temp Dir"`
		Version                      string `json:"Version"`
	} `json:"SonarQube"`
	Statistics struct {
		ID              string `json:"id"`
		Lines           int64  `json:"lines"`
		Ncloc           int64  `json:"ncloc"`
		NclocByLanguage []struct {
			Language string `json:"language"`
			Ncloc    int64  `json:"ncloc"`
		} `json:"nclocByLanguage"`
		Plugins []struct {
			Name    string `json:"name"`
			Version string `json:"version"`
		} `json:"plugins"`
		ProjectCount           int64 `json:"projectCount"`
		ProjectCountByLanguage []struct {
			Count    int64  `json:"count"`
			Language string `json:"language"`
		} `json:"projectCountByLanguage"`
		UserCount int64  `json:"userCount"`
		Version   string `json:"version"`
	} `json:"Statistics"`
	System struct {
		BootClassPath        string `json:"BootClassPath"`
		Daemon_Thread        int64  `json:"Daemon Thread"`
		Free_Memory          string `json:"Free Memory"`
		Heap                 string `json:"Heap"`
		JVM_Name             string `json:"JVM Name"`
		JVM_Vendor           string `json:"JVM Vendor"`
		JVM_Version          string `json:"JVM Version"`
		Library_Path         string `json:"Library Path"`
		Loaded_Classes       int64  `json:"Loaded Classes"`
		Max_Memory           string `json:"Max Memory"`
		Non_Heap             string `json:"Non Heap"`
		Processors           int64  `json:"Processors"`
		Start_Time           string `json:"Start Time"`
		System_Classpath     string `json:"System Classpath"`
		System_Date          string `json:"System Date"`
		System_Load_Average  string `json:"System Load Average"`
		Threads              int64  `json:"Threads"`
		Threads_Peak         int64  `json:"Threads Peak"`
		Total_Loaded_Classes int64  `json:"Total Loaded Classes"`
		Total_Memory         string `json:"Total Memory"`
		Unloaded_Classes     int64  `json:"Unloaded Classes"`
	} `json:"System"`
}
