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
            console.log("Player Added");
        },
        error: (d) => {
            console.log("An error occurred. Please try again");
        }
    });
}

function generateGames() {
    $.ajax({
        url: '/createGames',
        method: 'post',
        success: (d) => {
            console.log("Generated Games");
        },
        error: (d) => {
            console.log("An error occurred. Please try again");
        }
    });
}
