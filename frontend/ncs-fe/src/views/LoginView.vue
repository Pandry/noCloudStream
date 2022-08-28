<script>
import {reactive, ref} from 'vue'
import { useStore } from 'vuex'
import LoginTextbox from '@/components/LoginTextbox.vue'
import LoginSubmitButton from '@/components/LoginSubmitButton.vue'

/*const state = reactive({
    "username": "",
    "password": ""
})*/

/*export default {
  // `setup` is a special hook dedicated for composition API.
  setup() {
    const state = reactive({ count: 0 })

    // expose the state to the template
    return {
      state
    }
  }
}*/
/*const state = reactive({mail: "", password: ""})

function login(){
  alert("Attempting login with credentials" + `${state.mail}:${state.password}`)
  fetch('localhost:8080/api/auth/login', {
  method: 'POST',
  body: {"mail": state.mail, "password": state.password }}).then(response => {
    if (response.ok && response.status == 200){
      alert("login succeded!")
    }
     const contentType = response.headers.get('content-type');
     if (!contentType || !contentType.includes('application/json')) {
       throw new TypeError("Oops, we haven't got JSON!");
     }
     return response.json();
  })
  .then(data => {
      console.log(data)
  })
  .catch(error => console.error(error));

};*/
  export default {
    setup() {
      const state = reactive({mail: "", password: ""})

      const store = useStore()
      // expose the state to the template
      return {
        state,
        loginMutation: (s) => store.commit('login', s)
      }
    },
    components:{
      LoginTextbox,
      LoginSubmitButton
    },
    methods: {
        login() {
            // Authenticate user against API
            console.log(this.mail, this.password);
            
            // Set value somewhere to indicate that the user is logged in
            let isLoggedIn = true
            this.loginMutation(isLoggedIn)

            /*fetch('localhost:8080/api/auth/login', {
              method: 'POST',
              body: {"mail": state.mail, "password": state.password }}).then(response => {
                if (response.ok && response.status == 200){
                  alert("login succeded!")
                }
                const contentType = response.headers.get('content-type');
                if (!contentType || !contentType.includes('application/json')) {
                  throw new TypeError("Oops, we haven't got JSON!");
                }
                return response.json();
              })
              .then(data => {
                  console.log(data)
              })
              .catch(error => console.error(error));*/
            
            if (isLoggedIn) {
            // Redirect to page
              const redirectPath = this.$route.query.redirect || '/';
              this.$router.push(redirectPath);
            }
        },
    }
  }

</script>

<template>
  <div class="login">
    <h1 class="login-banner">Enter your credentials and login!</h1>
      <LoginTextbox v-model="state.mail" placeholder="Email" type="email"/>
      <LoginTextbox v-model="state.password" placeholder="Password" type="password"/>
      <LoginSubmitButton value="Log in" @click="login"/>
  </div>
</template>

<style>
@media (min-width: 1024px) {
  .login {
    min-height: 100vh;
    display: flex;
    align-items: center;
  }
}

.login{
  padding-top: 30vh;
  display: block
}

.login-banner{
  text-align: center;
}
</style>
