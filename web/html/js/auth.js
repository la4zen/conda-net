
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
                        localStorage.setItem("token", response.data.token)
                        location.href = "http://la4z.xyz"
                    }).catch(err => {
                        console.log(err)
                        alert(err.response.data)
                    })
                } else {
                    alert("Заполните все поля!")
                }
            },
            auth : function () {
                if (this.login && this.password) {
                    axios({
                        method:"POST",
                        headers: {
                            'Content-Type': 'application/json'
                        },
                        url : "http://la4z.xyz:8000/api/login",
                        data: JSON.stringify({
                            login:this.login,
                            password:this.password
                        })
                    }).then(response => {
                        localStorage.setItem("token", response.data.token)
                        location.href = "http://la4z.xyz"
                    }).catch(err => {
                        alert(err.response.data)
                    })
                } else {
                    alert("Заполните все поля!")
                }
            }
        }
    }).mount("body")
}