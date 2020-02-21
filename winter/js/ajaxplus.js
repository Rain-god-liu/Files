function ajax(obj){
	let defaults = {
		type : 'get',
		data : {},
		url : '#',
		dataType : 'json',
		async: true,
		success:function(data){console.log(data)},
		error:function(xhr){alert("请求发生错误或无资源查询"+xhr.status)}

	}
	for(let key in obj){
		defaults[key] = obj[key];
	}
	let xhr = null;
	if(window.XMLHttpRequest)
	{
		xhr = new XMLHttpRequest();
	}else {
		xhr = new ActiveXobject('Microsoft.XMLHTTP');
	}
	let param = '';
	for(let attr in obj.data){
		param += attr +'='+ obj.data[attr] + '&';
	}
	if(param){
		 param = param.substring(0,param.length-1);
	}
	if(defaults.type == 'get'){
		defaults.url +='?'+encodeURI(param);
	}
	xhr.open(defaults.type, defaults.url, defaults.async);
	let data = null;
	if(defaults.type =='post'){
		data = param;
		xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
	}
	xhr.send(data);
	xhr.onreadystatechange = function(){
		if(xhr.readyState == 4){
			if(xhr.status>=200&&xhr.status<300||xhr.status===304){
				let data = xhr.responseText;
				if(defaults.dataType == 'json'){
					data = JSON.parse(data);
				}
				defaults.success(data);
			}else {
				defaults.error(xhr)
			}
		}
	}
}

