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
            alert("An error occured. Please try again");
        }
    });
}
