<script setup>
import Login from './view/Login.vue';
import Landing from './view/Landing.vue';
</script>

<template>
	<!--<img id="logo" alt="Wails logo" src="./assets/images/logo-universal.png"/>-->
	<router-view></router-view>
	<!--<Landing />-->
</template>

<style>
#logo {
	display: block;
	width: 50%;
	height: 50%;
	margin: auto;
	padding: 10% 0 0;
	background-position: center;
	background-repeat: no-repeat;
	background-size: 100% 100%;
	background-origin: content-box;
}
</style>
