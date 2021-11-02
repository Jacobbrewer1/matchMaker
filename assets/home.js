function submitPlayer(fname, lname) {
    $.ajax({
        url: '/test',
        method: 'post',
        data: {
            fname: fname,
            lname: lname,
        },
        success: (d) => {
            alert("Player Added");
        },
        error: (d) => {
            alert("An error occured. Please try again");
        }
    });
}