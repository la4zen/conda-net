
window.onload = function () {
    Vue.createApp({
        data() {
            return {
                login : null,
                password : null,
                first_name : null,
                last_name : null,
                confirmPassword : null,
                user : null
            }
        },
        methods : {
            register : function() {
                if (this.password != this.confirmPassword) {
                    alert("Неверное подтверждение пароля")
                    return 
                }
                if (this.login && this.password && this.first_name && this.last_name) {
                    axios({
                        method:"POST",
                        headers: {
                            'Content-Type': 'application/json'
                        },
                        url : "http://la4z.xyz:8000/api/register",
                        data: JSON.stringify({
                            login:this.login,
                            password:this.password,
                            first_name : this.first_name,
                            last_name : this.last_name
                    })}).then((response) => {
                            if (response.status != 200) {
                                alert("Error")
                            } else {
                                localStorage.setItem("token", response.data.token)
                                localStorage.setItem("user", response.data.user)
                                location.href = "http://la4z.xyz"
                            }
                        })
                } else {
                    alert("Заполните все поля!")
                }
            },
            auth : function () {
                if (this.login && this.password && this.first_name && this.last_name) {
                    axios({
                        method:"POST",
                        headers: {
                            'Content-Type': 'application/json'
                        },
                        url : "http://la4z.xyz:8000/api/login",
                        json:{
                            login:this.login,
                            password:this.password
                        }
                    }).then((response) => {
                        if (response.code != 200) {
                            alert("Error")
                        } else {
                            localStorage.setItem("token", response.data.token)
                            localStorage.setItem("user", response.data.user)
                            location.href = "http://la4z.xyz"
                        }
                    })
                } else {
                    alert("Заполните все поля!")
                }
            }
        }
    }).mount("body")
}