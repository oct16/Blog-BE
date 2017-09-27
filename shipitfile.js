module.exports = function (shipit) {
  require('/Users/oct16/.config/yarn/global/node_modules/shipit-deploy')(shipit)

  shipit.initConfig({
    default: {
      workspace: '/tmp/blog_be',
      deployTo: '/home/blog_be/echo-blog',
      repositoryUrl: 'https://github.com/oct16/Blog-BE.git',
      branch: 'master',
      ignores: ['.gitignore', 'shipitfile.js'],
      rsync: ['--del'],
      keepReleases: 1,
      key: '~/.ssh/id_rsa',
      shallowClone: true,
      servers: 'root@97.64.19.213:27471'
    }
  })

  shipit.on('published', function () {
      return shipit.start(['copyFiles', /*'copyDepends', */'cleanContainer', /*'buildImage',*/ 'run'])
  })

  shipit.blTask('copyFiles', function() {
    return shipit.remote(`
      cd ${shipit.currentPath} &&
      mkdir -p /home/go_files/echo-blog &&
      rm -rf /home/go_files/echo-blog/* &&
      cp -r ./* /home/go_files/echo-blog
    `)
  })

  shipit.blTask('copyDepends', function() {
    return shipit.remote(`cp -r /home/depends /home/go_files/echo-blog/depends`)
  })

  shipit.blTask('cleanContainer', function() {
    return shipit.remote(`docker stop blog_be || true && docker rm blog_be || true && docker rmi -f blog_be || true`)
  })

  shipit.blTask('buildImage', function() {
    return shipit.remote(`cd /home/go_files/echo-blog/ && docker build -t blog_be .`)
  })

  shipit.blTask('run', function() {
    return shipit.remote(`docker run --name blog_be -v /home/go_files:/go/src -d -p 3016:3016 golang /bin/bash -c "cd /go/src/echo-blog && ./gorun.sh"`)
  })
}
