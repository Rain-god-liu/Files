let inputs = document.querySelectorAll('input');
for(let i =0;i<=1;i++){
	inputs[i].addEventListener('focus',function(){
		if(i==0){
			this.placeholder = '手机号';
			this.classList.remove('change');
		}else{
			this.placeholder = '输入6位短信验证码';
			this.classList.remove('change');
		}
	})
	inputs[i].addEventListener('blur',function(){
		if(i==0){
			this.placeholder = '请输入手机号';
			this.className = 'change';
		}else{
			this.placeholder = '请输入短信验证码';
			this.className = 'change';
		}
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
	})
}