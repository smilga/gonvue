<template>
    <div class="app">
        <div class="nav">
            <div class="links">
                <router-link to="/">Home</router-link>
                <router-link to="/about">About</router-link>
            </div>
        </div>
        <router-view></router-view>
    </div>
</template>

<style lang="scss">
@import url('https://fonts.googleapis.com/css?family=Varela+Round&display=swap');
* {
    font-family: 'Varela Round', sans-serif;
    color: #3c3c3c;
}

.nav {
    display: flex;
    align-items: center;
    justify-content: center;
    margin-top: 20px;
}
</style>
