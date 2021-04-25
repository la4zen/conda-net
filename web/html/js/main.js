window.onload = function() {
    Vue.createApp({
        data() {
            return {
                user: {
                    login : null,
                    first_name : null,
                    last_name : null,
                    last_login : null
                }
            }
        },
        created() {
            if (!localStorage.getItem("token")) {
                location.href = "http://la4z.xyz/login"
                return 
            }
            axios({
                method:"POST",
                url : "http://la4z.xyz:8000/api/users/get",
                headers : {
                    "Authorization" : "Bearer " + localStorage.getItem("token")
                }
            }).then((response) => {
                if (response.status != 200) {
                    location.href = "http://la4z.xyz/login"
                } else {
                    this.user = response.data.user
                }
            })
        },
        methods : {
            getTime : function () {
                return moment(Date.parse(this.user.last_login)).locale("ru").fromNow()
            }
        }
    }).mount("body")
}