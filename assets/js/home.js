var players = document.querySelector('#players');
var gamesDisplayer = document.querySelector('#games')

function onSubmit(evt) {
    evt.preventDefault()
    let form = evt.target;
    let formData = new FormData(form);
    let fname = formData.get('firstName');
    let lname = formData.get('lastName');
    let gender = formData.get('gender');
    if (fname != "" || lname != "" || gender != "") {
        $('#id01').hide();
        $('#playerSpinner').show();
        $.ajax({
            url: '/addPlayer',
            method: 'post',
            data: formData,
            processData: false,
            contentType: false,
            success: (d) => {
                $('#playerSpinner').hide();
                console.log("Player Added", d);
                players.innerHTML += d;
                form.reset();
            },
            error: (d) => {
                $('#playerSpinner').hide();
                console.log("An error occurred. Please try again");
            }
        });
    }

    return false;
}

function generateGames() {
    $.ajax({
        url: '/createGames',
        method: 'post',
        success: (d) => {
            document.getElementById('generateGames').classList.toggle('button--loading')
            console.log("Games generated", d);
            gamesDisplayer.innerHTML += d;
        },
        error: (d) => {
            document.getElementById('generateGames').classList.toggle('button--loading')
            console.log("An error occurred. Please try again");
        }
    });
}

function clearBackend() {
    $.ajax({
        url: '/cleanse',
        method: 'post',
        success: (d) => {
            document.getElementById('cleanseProgramButton').classList.toggle('button--loading')
            console.log("Backend cleared", d);
            gamesDisplayer.innerHTML += d;
        },
        error: (d) => {
            document.getElementById('cleanseProgramButton').classList.toggle('button--loading')
            console.log("An error occurred. Please try again");
        }
    });
}
