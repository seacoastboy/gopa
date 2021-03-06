/** 
 * User: Medcl
 * Date: 13-7-23
 * Time: 下午2:20 
 */
package config
 import (
	 config "github.com/robfig/config"
	 log "logging"
 )

var loadingConfig *config.Config

func  init(){
	//parse main config
	loadingConfig, _ = config.ReadDefault("config.ini")
}

func GetStringConfig(configSection string,configKey string ,defaultValue string) string{
	if(loadingConfig ==nil){
		log.Trace("loadingConfig is nil,just return")
		return defaultValue
	}

	//loading or initializing bloom filter
	value,error:=loadingConfig.String(configSection, configKey)
	if(error!=nil){
		value=defaultValue
	}
	log.Trace("get config value,",configSection,".",configKey,":",value)
	return value
}
func GetFloatConfig(configSection string,configKey string ,defaultValue float64) float64{
	if(loadingConfig ==nil){
		log.Trace("loadingConfig is nil,just return")
		return defaultValue
	}

	//loading or initializing bloom filter
	value,error:=loadingConfig.Float(configSection, configKey)
	if(error!=nil){
		value=defaultValue
	}
	log.Trace("get config value,",configSection,".",configKey,":",value)
	return value
}
func GetIntConfig(configSection string,configKey string ,defaultValue int) int{
	if(loadingConfig ==nil){
		log.Trace("loadingConfig is nil,just return")
		return defaultValue
	}

	//loading or initializing bloom filter
	value,error:=loadingConfig.Int(configSection, configKey)
	if(error!=nil){
		value=defaultValue
	}
	log.Trace("get config value,",configSection,".",configKey,":",value)
	return value
}

func GetBoolConfig(configSection string,configKey string ,defaultValue bool) bool{
	if(loadingConfig ==nil){
		log.Trace("loadingConfig is nil,just return")
		return defaultValue
	}

	//loading or initializing bloom filter
	value,error:=loadingConfig.Bool(configSection, configKey)
	if(error!=nil){
		value=defaultValue
	}
	log.Trace("get config value,",configSection,".",configKey,":",value)
	return value
}

