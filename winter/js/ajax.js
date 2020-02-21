//url?key=value&key=value;
/*function obj2 (obj){
    var  str1= [];
    for(var key in obj){
        str1.push(encodeURIComponent(key)+'='+encodeURIComponent(obj[key]));
    } 
    return str1.join("&")
}*/
function ajax(url,type,obj,datatype,success,error){
    //0.将对象转换成字符串
    //var str = obj2(obj);
    //1.申请一个异步对象
    var xmlhttp = new XMLHttpRequest();
    //1.5判断请求类型
    if(type=='get'){ //get请求
        url += '?key='+obj;
    }
    let data = null;
    if(type == 'post'){ //post请求
        data = obj;
        xmlhttp.setRequestHeader("Content-Type", "application/x-www-form-urlencoded")
    }
    //2.设置请求方式和地址
    xmlhttp.open("GET",url,true); 
    //3.发送请求
    console.log(xmlhttp.status);
        xmlhttp.send(data);
    //4.监听状态变化
    xmlhttp.onreadystatechange = function(){
        if(xmlhttp.readyState===4){
            if(xmlhttp.status>=200&&xmlhttp.status<300||xmlhttp.status===304){
                let result = xmlhttp.responseText;
                if(result == 'json'){
                    result = JSON.parse(result);
                }
                success(result);
            }else{
                error(xmlhttp);
                }
        }    
    }   
    
}