const avatarContainer = document.getElementById('avatarContainer');

const avatars = ['images/avatars/avatar1.png',
                'images/avatars/avatar2.png',
                'images/avatars/avatar3.png',
                'images/avatars/avatar4.png',
                'images/avatars/avatar5.png',
                'images/avatars/avatar6.png',
                'images/avatars/avatar7.png',
                'images/avatars/avatar8.png',
                'images/avatars/avatar9.png',
                'images/avatars/avatar10.png',
                'images/avatars/avatar11.png',
                'images/avatars/avatar12.png',
                'images/avatars/avatar13.png',
                'images/avatars/avatar14.png',
                'images/avatars/avatar15.png',
];

function generateAvatars(){
    avatars.forEach((avatars,index)=>{
        const avatarDiv = document.createElement('div');
        avatarDiv.classList.add('avatar-item');
        avatarDiv.innerHTML = `<img src ="${avatars}" alt = Avatar ${index+1} />`;

        avatarDiv.addEventListener('click',()=>{
            document.querySelectorAll('.avatar-item').forEach(item => item.classList.remove('selected'));
            avatarDiv.classList.add('selected');

        });
        avatarContainer.appendChild(avatarDiv);
        
    });
    
}
generateAvatars();