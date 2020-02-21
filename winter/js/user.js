window.addEventListener('load',function(){
	let fllowbutton = document.querySelector('.Follow-button');
	fllowbutton.addEventListener('mousemove',function(){
		fllowbutton.innerHTML = '取消关注';

	})
	fllowbutton.addEventListener('mouseout',function(){
		fllowbutton.innerHTML = '已关注';
	});
	change = document.querySelector('#change');
	let ProfileHeadertitle = document.querySelector('.ProfileHeader-title');
	let Fieldmodify = document.querySelectorAll('.Field-modify');
	let Fieldcontent = document.querySelector('.Field-content');
	console.log(Fieldcontent);
	let fieldchoice = Fieldcontent.children;
	console.log(fieldchoice);
	Fieldcontent.addEventListener('mousemove',function(){
		Fieldmodify[1].classList.remove('Field-modify-hidden');
	});
	Fieldcontent.addEventListener('mouseout',function(){
		Fieldmodify[1].classList.add('Field-modify-hidden');
	});
	ProfileHeadertitle.addEventListener('mousemove',function(){
		Fieldmodify[0].classList.remove('Field-modify-hidden');
	});
	ProfileHeadertitle.addEventListener('mouseout',function(){
		Fieldmodify[0].classList.add('Field-modify-hidden');
	});
	Fieldmodify[1].addEventListener('click', function(){    		//“修改”按钮的点击事件 修改模块和显示模块的切换
		fieldchoice[0].style.display = 'none';
		fieldchoice[1].style.display = 'block';
	})
	let allbutton = document.querySelectorAll('.allbutton');  //得到“保存”和“取消”按钮
	allbutton[0].addEventListener('click',function(){   //保存按钮的点击事件
		fieldchoice[1].style.display = 'none';
		fieldchoice[0].style.display = 'block';
	})
	/*allbutton[1].addEventListener('click',function(){  //取消按钮的点击事件
		for(let i = 0;i<sexinput.length;i++){
	
		}
	})*/
	change.addEventListener('click',function(){
		Profilemain = document.querySelector('.Profile-main');
		ProfileHeadercontentBody = document.querySelector('.ProfileHeader-contentBody');
		ProfileHeadercontentFooter = document.querySelector('.ProfileHeader-contentFooter');
		ProfileHeadercontentFooter.style.display = 'none';
		ProfileHeadercontentBody.style.display = 'none';
		Profilemain.style.display = 'none';
	})
})
