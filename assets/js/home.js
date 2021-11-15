var players = document.querySelector('#players');

function onSubmit(evt) {
    evt.preventDefault()
    let form = evt.target;
    let formData = new FormData(form);
    let fname = formData.get('firstName');
    let lname = formData.get('lastName');
    let gender = formData.get('gender');
    if (fname != "" || lname != "" || gender != "") {
        $('#playerSpinner').show();
        $('#id01').hide();
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
            console.log("Generated Games");
            location.reload();
        },
        error: (d) => {
            console.log("An error occurred. Please try again");
        }
    });
}
