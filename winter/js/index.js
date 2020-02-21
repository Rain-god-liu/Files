let userimg = document.getElementById('user-img');
let contentMune = document.querySelector('.content-Mune');
let flag = 0;
let flags = 0;
let VoteButton = document.querySelectorAll('.VoteButton');
let RichContentbutton = document.querySelector('.RichContent-button')
//let tabs = document.querySelector('.tabs');
let richcontent = document.querySelector('.RichContent-innner');
let ContentItemactions = document.querySelector('.ContentItem-actions'); 
let items = document.querySelectorAll('.tabs-items');
let divs = document.querySelector('.ListShortcut-choice').children;
for(var i = 0;i<items.length;i++){
	items[i].setAttribute('data-index', i);
	items[i].addEventListener('click',function(){
		for(let i = 0;i<divs.length;i++){
			divs[i].style.display = 'none';
			items[i].classList.remove('link-blue');
		}
		this.classList.add('link-blue');
		let index = this.getAttribute('data-index');
		divs[index].style.display = 'block';
	})	
};
userimg.addEventListener('click',function(){
	if(flag==0){
		contentMune.style.display = 'block';
		flag = 1;
	}else {
		contentMune.style.display = 'none';
		flag = 0;
	}
});
richcontent.addEventListener('click',function(){
		richcontent.classList.remove('qusetionshow');
		ContentItemactions.style.position = 'fixed';
		RichContentbutton.style.display = 'none';
		ContentItemactions.classList.add('activc-click');
});
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

