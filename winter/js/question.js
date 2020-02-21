let anwser = document.querySelector('#Anwser');
let anwsershow = document.querySelector('.ViewAll').children;
let flag = 0;
let flags = 0;
let QusetionShow = document.querySelector('#QusetionShow');
let VoteButton = document.querySelectorAll('.VoteButton')
let QuestionHeaderactions = document.querySelector('.QuestionHeader-actions');
let Qusetion = QusetionShow.children;
let questiondescipt = document.querySelector('#question-value');
let questionvalue = questiondescipt.innerHTML;
let p = document.createElement('p');
let button = document.createElement('button');
console.log(questionvalue);
anwser.addEventListener('click', function(){
	if(flag == 0){
		anwsershow[0].style.display = 'none';
		anwsershow[1].style.display = 'block';
		flag = 1;
	}else {
		anwsershow[1].style.display = 'none';
		anwsershow[0].style.display = 'block';
		flag = 0;
	}
});
//1.打开全文
QusetionShow.addEventListener('click',function(){
	QusetionShow.parentElement.classList.remove('QuestionRichText');
	let showbutton = document.querySelector('#Show-button');
	//showbutton.style.display = 'none';
	QusetionShow.removeChild(showbutton);
	QuestionHeaderactions.style.display = 'block';
	questiondescipt.innerHTML = '';
	QusetionShow.children[0].appendChild(p);
	p.innerHTML = questionvalue;
});
//2.收起全文
QuestionHeaderactions.addEventListener('click',function(){
	QusetionShow.parentElement.classList.add('QuestionRichText');
	//showbutton.style.display = 'block';

	QusetionShow.appendChild(button);
	button.className= 'Button--Grey';
	button.id = 'Show-button';
	button.innerHTML = '显示全部'+'<svg viewBox="0 0 10 6" class="Icon QuestionRichText-more-icon Icon--arrow" width="10" height="16" aria-hidden="true" style="height: 16px; width: 10px;"><title></title><g><path d="M8.716.217L5.002 4 1.285.218C.99-.072.514-.072.22.218c-.294.29-.294.76 0 1.052l4.25 4.512c.292.29.77.29 1.063 0L9.78 1.27c.293-.29.293-.76 0-1.052-.295-.29-.77-.29-1.063 0z"></path></g></svg>';
	QuestionHeaderactions.style.display = 'none';
	QusetionShow.children[0].removeChild(p);
	questiondescipt.innerHTML = questionvalue;
});
//3.添加赞同 差评的点击样式
for (let i = 0; i < VoteButton.length; i++) {
	VoteButton[i].addEventListener('click',function(){
		if(flags==0){
			VoteButton[i].classList.add('active');
			VoteButton[i].classList.remove('VoteButton');
			VoteButton[0].innerHTML = '已赞同';
			flags = 1;
		}else{
			VoteButton[i].classList.remove('active');
			VoteButton[i].classList.add('VoteButton');
			flags = 0;	
			VoteButton[0].innerHTML ='<svg class="" fill="currentColor" viewBox="0 0 24 24" width="10" height="10"><path d="M2 18.242c0-.326.088-.532.237-.896l7.98-13.203C10.572 3.57 11.086 3 12 3c.915 0 1.429.571 1.784 1.143l7.98 13.203c.15.364.236.57.236.896 0 1.386-.875 1.9-1.955 1.9H3.955c-1.08 0-1.955-.517-1.955-1.9z" fill-rule="evenodd"></path></svg>'+'赞同'
			
		}
		
	})
}
