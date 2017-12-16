$(function () {
    $(".gq-anwser-btn").click(function () {
        $(this).html($(this).html() == "隐藏答案" ? "显示答案" : "隐藏答案").next(".gq-anwser").toggle()
    })
    $(".gq-run-btn").click(function () {
        var $this = $(this), $code, $output
        if ($this.prev().is(".output")) {
            $output = $this.prev()
            $code = $output.prev().find("code")
            $output.removeClass().addClass("output bs-callout").html("<img src='/assets/compiling.gif' />")
        } else {
            $code = $this.prev().find("code")
            $this.before("<div class='output bs-callout'><img src='/assets/compiling.gif' /></div>")
            $output = $this.prev()
        }
        $.ajax({
            type: "POST",
            headers: {
                "x-requested-with": location.host
            },
            dataType: "json",
            data: { body: $('<textarea />').html($code.html()).text() },
            url: "https://cors-anywhere.herokuapp.com/https://play.golang.org/compile",
            success: function (data) {// bs-callout-warning bs-callout-danger bs-callout-info
                var output, cls
                if(data.compile_errors){
                    output = data.compile_errors
                    cls = "bs-callout-danger"
                }else if(data.Errors){
                    output = data.Errors
                    cls = "bs-callout-warning"
                }else{
                    if(data.output){
                        output = data.output
                    }else {
                        output = $.map(data.Events,e=>e.Message).join("<br/>")
                    }
                    cls = "bs-callout-info"
                }
                $output.addClass(cls).html(output.replace(/\n/g, "<br/>"))
            }
        });
    })
})