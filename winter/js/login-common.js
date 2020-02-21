let div = document.querySelector('.header');
let lione = document.querySelector('#one');
let spans = document.querySelector('#two').querySelectorAll('span');
let SignFlowsupportedCount = document.querySelector('.SignFlow-supportedCount');
let SignFlowaccountSeperator = document.querySelector('.SignFlow-accountSeperator');
let scripts = document.querySelector('script');
let input = document.querySelectorAll('input');
let user = lione.children;
let divs = div.children;
let tlp = document.querySelector('#th').children;
let fobutton = document.querySelector('#fo').querySelector('button');
console.log(fobutton);
fobutton.addEventListener('click',function(){
	console.log(input[0].value);
	console.log(input[1].value);
	ajax({
	type:'post',
	data:{username:input[0].value,password:input[1].value},
	url:'http://localhost:8080/register',
	success:function(data){
		if(data.status == 200){
			ajax({
				type:'post',
				data:{username:input[0].value,password:input[1].value},
				url:'http://localhost:8080/login',
				success:function(data){
					console.log(data);
				}
			})
		}
	}
	});
});
	
for(let i = 0;i<=1;i++){
	divs[i].addEventListener('click',function(){
		if(i==0){
			divs[0].classList.add('border');
			divs[1].classList.remove('border');
			SignFlowsupportedCount.style.display = 'block';
			SignFlowaccountSeperator.style.display = 'block';
			input[0].placeholder = '手机号';
			input[1].placeholder = '输入6位短信验证码';
			tlp[1].innerHTML = '接受语音短信验证码';
			for(let i = 0;i<input.length;i++){
				input[i].classList.remove('change');
			}
			
			for(let j = 0;j<spans.length;j++){
				spans[j].innerHTML = '获取短信验证码';
				spans[j].className = 'Verification';
			}
			for(let i =0;i<=1;i++){
				input[i].addEventListener('focus',function(){
					if(i==0){
						this.placeholder = '手机号';
						this.classList.remove('change');
					}else{
						this.placeholder = '输入6位短信验证码';
						this.classList.remove('change');
					}
				})
				input[i].addEventListener('blur',function(){
					if(i==0){
						this.placeholder = '请输入手机号';
						this.className = 'change';
					}else{
						this.placeholder = '请输入短信验证码';
						this.className = 'change';
					}
				})
					let lis = document.querySelector('#two');
					eyesbutton = lis.children;
					for(let i = 1;i<=2;i++){
						eyesbutton[i].addEventListener('click',function(){
							if(i==2){
								eyesbutton[2].style.display = 'none';
								eyesbutton[1].style.display = 'block';
								eyesbutton[0].type = 'text';
							}else {
								eyesbutton[1].style.display = 'none';
								eyesbutton[2].style.display = 'block';
								eyesbutton[0].type = 'password';
								}
						});
					}
			}

		}else {
			divs[1].classList.add('border');
			divs[0].classList.remove('border');
			spans[0].innerHTML = '';
			spans[1].innerHTML  ='';
			input[0].placeholder = '手机号或者邮箱';
			input[1].placeholder = '密码';
			SignFlowsupportedCount.style.display = 'none';
			SignFlowaccountSeperator.style.display = 'none';
			let lis = document.querySelector('#two');
			tlp[1].innerHTML = '忘记密码?';
			eyesbutton = lis.children;
			for(let i = 0;i<input.length;i++){
				input[i].classList.remove('change');
			}
			for(let i = 1;i<=2;i++){
			eyesbutton[i].addEventListener('click',function(){
				if(i==2){
					eyesbutton[2].style.display = 'none';
					eyesbutton[1].style.display = 'block';
					eyesbutton[0].type = 'text';
				}else {
					eyesbutton[1].style.display = 'none';
					eyesbutton[2].style.display = 'block';
					eyesbutton[0].type = 'password';
					}
				});
			}
			for(let i = 0;i<input.length;i++){
				input[i].addEventListener('focus',function(){
					if(i==0){
						this.placeholder = '手机号或者邮箱';
						this.classList.remove('change');
					}else{
						this.placeholder = '密码';
						this.classList.remove('change');
					}
				})
				input[i].addEventListener('blur',function(){
					if(i==0){
						this.placeholder = '请输入手机号或者邮箱';
						this.className = 'change';
					}else{
						this.placeholder = '请输入密码';
						this.className = 'change';
					}
				})
			}
			for(let l = 0;l<spans.length;l++){
				spans[l].className = 'button-eyes';
			}
		}
	})
};