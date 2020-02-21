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
let inputs = document.querySelectorAll('input');
for(let i =0;i<=1;i++){
	inputs[i].addEventListener('focus',function(){
		if(i==0){
			this.placeholder = '手机号或者邮箱';
			this.classList.remove('change');
		}else{
			this.placeholder = '密码';
			this.classList.remove('change');
		}
	})
	inputs[i].addEventListener('blur',function(){
		if(i==0){
			this.placeholder = '请输入手机号或邮箱';
			this.className = 'change';
		}else{
			this.placeholder = '请输入密码';
			this.className = 'change';
		}
	})
}