$(document).ready(function () {

    $("#LoginEntity_Captcha").val("0");

    $('#form2').hide();
    $('#lnkForgotPassword').click(function () {
        $('#form1').hide('blind', 1000, $('#form2').show());
    });

    $('#btnCancelReset').click(function () {
        $('#form2').hide();
        $('#form1').show()
    });

    //uncomment this for invisible capta    
    //$("#btnLogin").click(function () {
    //    grecaptcha.execute();
    //});
});

var CaptchaWidget;
var verifyCallback = function (response) {   
    if (response != null && response != "") {
        $("#LoginEntity_Captcha").val(response);
        return true;
    }
    else {
        $("#LoginEntity_Captcha").val("");
        return false;
    }
};

var ExpiredCallback = function () {
    $("#LoginEntity_Captcha").val("");
    return false;
};

var onloadCallback = function () {
    // Renders the HTML element with id 'divRecaptcha' as a reCAPTCHA widget.
    // The id of the reCAPTCHA widget is assigned to 'CaptchaWidget'.
    
    //uncomment this for regular capta
    //CaptchaWidget = grecaptcha.render('divRecaptcha', {
    //    'sitekey': '6Leg2l8UAAAAALKhE2bBDav1U5ZLBG4iLjzb5U21',
    //    'theme': 'light',
    //    'callback': verifyCallback,
    //    'expired-callback': ExpiredCallback
    //});
    
    //uncomment this for invisible capta
    //Capta = grecaptcha.render('btnLogin', {
    //    'sitekey': '6Lexx2MUAAAAAJaH1aM1eLVzrKlzZbx-4vcIZa5X',
    //    'theme': 'light',
    //    'callback': onSubmit
    //});
}

var slideIndex = 0;
carousel();

function carousel() {
    var i;
    var x = document.getElementsByClassName("mySlides");
    for (i = 0; i < x.length; i++) {
        x[i].style.display = "none";
    }
    slideIndex++;
    if (slideIndex > x.length) { slideIndex = 1 }
    x[slideIndex - 1].style.display = "block";
    setTimeout(carousel, 5000); // Change image every 2 seconds
}

function onSubmit(token) {
    
    $("#LoginEntity_Captcha").val(token);
    $('#frmLogin').submit();
    ResetCapta();
   
}
function ResetCapta() {
    $("#LoginEntity_Captcha").val("");
    grecaptcha.reset();
}