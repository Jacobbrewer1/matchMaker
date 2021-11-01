function submitPlayer() {
    $.ajax({
        url: '/test',
        method: 'post',
        success: (d) => {
            console.log("Success")
        },
        error: (d) => {
            console.log("Error")
        }
    });
}