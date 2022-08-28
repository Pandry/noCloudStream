<template>
  <div class="signup">
    <h1 class="signup-banner">Register to StreamField!</h1>
    <LoginTextbox v-model="state.name" placeholder="Your name" type="text"/>
    <LoginTextbox v-model="state.mail" placeholder="Your mail" type="email"/>
    <LoginTextbox v-model="state.password" placeholder="Your password" type="password"/>
    <LoginSubmitButton value="Sign up" @click="signup"/>
  </div>
</template>

<script setup>
import {reactive} from 'vue'
import LoginTextbox from '@/components/LoginTextbox.vue'
import LoginSubmitButton from '@/components/LoginSubmitButton.vue'
const state = reactive({name:"", mail: "", password: ""})

function signup(e){
  e.preventDefault();
  fetch('localhost:8080/api/auth/signup', {
  method: 'POST',
  body: {"mail": state.mail, "password": state.password, name: state.name }}).then(response => {
    if (response.ok && response.status == 201){
      alert("registration succeded!")
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
}
</script>

<style>
@media (min-width: 1024px) {
  .signup {
    min-height: 100vh;
    display: flex;
    align-items: center;
  }
}

.signup{
  padding-top: 30vh;
  display: block
}

.signup-banner{
  text-align: center;
}
</style>
