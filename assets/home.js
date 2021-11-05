function submitPlayer(fname, lname, gender) {
    if (fname == "" || lname == "" || gender == "") {
        return;
    }
    $.ajax({
        url: '/addPlayer',
        method: 'post',
        data: {
            fname: fname,
            lname: lname,
            gender: gender,
        },
        success: (d) => {
            alert("Player Added");
        },
        error: (d) => {
            alert("An error occurred. Please try again");
        }
    });
}

function generateGames() {
    $.ajax({
        url: '/createGames',
        method: 'post',
        success: (d) => {
            alert("Generated Games");
        },
        error: (d) => {
            alert("An error occurred. Please try again");
        }
    });
}
