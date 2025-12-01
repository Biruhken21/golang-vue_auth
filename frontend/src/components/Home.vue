<template>
  <div class="min-h-screen flex flex-col items-center justify-center bg-gradient-to-br from-vt-c-bg to-vt-c-black-mute text-vt-c-text-1 py-16 px-4 text-center w-full">
    
    <!-- Outer container that controls max width -->
    <div class="w-full max-w-7xl px-4 sm:px-8 mx-auto">
      
      <!-- Content container -->
      <div class="home-container bg-vt-c-bg-soft p-8 sm:p-16 rounded-xl shadow-2xl flex flex-col items-center space-y-12">

        <!-- Header -->
        <div>
          <h1 class="text-4xl sm:text-5xl font-extrabold text-vt-c-heading leading-tight mb-6 tracking-tight">
            Welcome to Our Awesome App!
          </h1>
          <p class="text-lg sm:text-xl text-vt-c-text-2 leading-relaxed max-w-2xl mx-auto">
            Discover amazing features and streamline your daily tasks with ease. Get started now and experience the difference!
          </p>
        </div>

        <!-- Buttons -->
        <div class="flex flex-col sm:flex-row justify-center space-y-4 sm:space-y-0 sm:space-x-4 w-full">
          <router-link to="/dashboard" class="inline-flex justify-center items-center px-8 py-4 rounded-lg font-bold no-underline transition-all duration-300 ease-in-out bg-indigo-600 text-white shadow-md hover:bg-indigo-700 hover:shadow-lg transform hover:-translate-y-1 border border-transparent focus:outline-none focus:ring-4 focus:ring-indigo-500/40">
            Explore Features
          </router-link>
          <router-link to="/about" class="inline-flex justify-center items-center px-8 py-4 rounded-lg font-bold no-underline transition-all duration-300 ease-in-out bg-transparent text-indigo-600 border border-indigo-600 hover:bg-indigo-600 hover:text-white hover:border-indigo-700 transform hover:-translate-y-1 focus:outline-none focus:ring-4 focus:ring-indigo-500/30">
            Learn More
          </router-link>
        </div>

        <!-- Feature Boxes -->
        <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6 w-full mt-8">
          <div class="p-6 rounded-xl bg-black/20 border border-white/10 text-left shadow hover:shadow-lg transition">
            <div class="w-10 h-10 rounded-lg bg-indigo-500/20 text-indigo-300 flex items-center justify-center mb-3">
              <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" class="w-6 h-6">
                <path d="M12 6v12m6-6H6" />
              </svg>
            </div>
            <h3 class="text-xl font-semibold mb-2">Fast Performance</h3>
            <p class="text-vt-c-text-2">Optimized for speed so you can get things done quicker.</p>
          </div>

          <div class="p-6 rounded-xl bg-black/20 border border-white/10 text-left shadow hover:shadow-lg transition">
            <div class="w-10 h-10 rounded-lg bg-emerald-500/20 text-emerald-300 flex items-center justify-center mb-3">
              <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" class="w-6 h-6">
                <path fill-rule="evenodd" d="M2.25 12c0-5.385 4.365-9.75 9.75-9.75s9.75 4.365 9.75 9.75-4.365 9.75-9.75 9.75S2.25 17.385 2.25 12Zm13.36-2.59a.75.75 0 0 0-1.22-.87l-3.236 4.542-1.64-1.64a.75.75 0 1 0-1.06 1.06l2.25 2.25a.75.75 0 0 0 1.154-.114l3.812-5.228Z" clip-rule="evenodd" />
              </svg>
            </div>
            <h3 class="text-xl font-semibold mb-2">Secure by Default</h3>
            <p class="text-vt-c-text-2">Best practices and modern auth to keep your data safe.</p>
          </div>

          <div class="p-6 rounded-xl bg-black/20 border border-white/10 text-left shadow hover:shadow-lg transition">
            <div class="w-10 h-10 rounded-lg bg-purple-500/20 text-purple-300 flex items-center justify-center mb-3">
              <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" class="w-6 h-6">
                <path d="M3 4a1 1 0 0 1 1-1h5a1 1 0 0 1 .8.4l1.9 2.6H20a1 1 0 0 1 1 1v10.5a1.5 1.5 0 0 1-1.5 1.5H4.5A1.5 1.5 0 0 1 3 17.5V4Z" />
              </svg>
            </div>
            <h3 class="text-xl font-semibold mb-2">Beautiful UI</h3>
            <p class="text-vt-c-text-2">Carefully crafted design system with responsive layouts.</p>
          </div>
        </div>

        <!-- Testimonials -->
        <section class="w-full mt-16">
          <h2 class="text-2xl font-bold mb-6">What our users say</h2>
          <div class="grid grid-cols-1 md:grid-cols-3 gap-6 text-left">
            <div class="p-6 rounded-xl bg-black/20 border border-white/10 shadow">
              <p class="text-vt-c-text-2">“This app changed how I work every day. Super fast and reliable.”</p>
              <div class="mt-4 text-sm">— Alex, Product Manager</div>
            </div>
            <div class="p-6 rounded-xl bg-black/20 border border-white/10 shadow">
              <p class="text-vt-c-text-2">“Beautiful UI and easy to use. I love the attention to detail.”</p>
              <div class="mt-4 text-sm">— Sam, Designer</div>
            </div>
            <div class="p-6 rounded-xl bg-black/20 border border-white/10 shadow">
              <p class="text-vt-c-text-2">“Secure and dependable. Exactly what our team needed.”</p>
              <div class="mt-4 text-sm">— Jamie, Engineer</div>
            </div>
          </div>
        </section>

        <!-- CTA Section -->
        <section class="w-full mt-16">
          <div class="p-8 rounded-xl bg-indigo-600 text-white flex flex-col sm:flex-row items-center justify-between gap-4">
            <div class="text-left">
              <h3 class="text-2xl font-bold">Ready to get started?</h3>
              <p class="opacity-90">Join thousands of users boosting their productivity today.</p>
            </div>
            <router-link to="/register" class="inline-flex justify-center items-center px-6 py-3 rounded-lg font-semibold bg-white text-indigo-700 hover:bg-gray-100 transition">
              Create an account
            </router-link>
          </div>
        </section>

        <!-- Footer -->
        <footer class="mt-16 pt-6 border-t border-white/10 w-full text-sm text-vt-c-text-2">
          <p>© {{ new Date().getFullYear() }} Awesome App. All rights reserved.</p>
        </footer>

      </div>
    </div>
  </div>
</template>
