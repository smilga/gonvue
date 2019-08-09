<template>
    <div>
        <div class="title">
            <div class="love">
                <img class="gopher" src="../assets/img/gopher.png" alt="">
                <img class="heart" src="../assets/img/love.png" alt="">
                <img class="vue" src="../assets/img/vue.png" alt="">
            </div>
            <h1>Go Vue friendly template</h1>
            <h2>Start building something!</h2>
        </div>
    </div>
</template>

<script>
export default {
    name: 'Home'
}
</script>

<style lang="scss">
.title {
    display: flex;
    align-items: center;
    flex-direction: column;
}
.love {
    margin-top: 80px;
    margin-bottom: 40px;
    display: flex;
    align-items: center;
    img {
        max-height: 200px;
    }
    .heart {
        height: 50px;
        margin-left: 50px;
    }
}
</style>
