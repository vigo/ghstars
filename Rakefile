# frozen_string_literal: true

# constants
# -----------------------------------------------------------------------------
AVAILABLE_REVISIONS = %w[major minor patch].freeze
# -----------------------------------------------------------------------------


# -----------------------------------------------------------------------------
# hidden tasks
# -----------------------------------------------------------------------------
task :command_exists, [:command] do |_, args|
  abort "#{args.command} doesn't exists" if `command -v #{args.command} > /dev/null 2>&1 && echo $?`.chomp.empty?
end

task :repo_clean do
  abort 'please commit your changes first!' unless `git status -s | wc -l`.strip.to_i.zero?
end

task :has_bumpversion do
  Rake::Task['command_exists'].invoke('bumpversion')
end

task :has_gsed do
  Rake::Task['command_exists'].invoke('gsed')
end

task :bump, [:revision] => [:has_bumpversion] do |_, args|
  args.with_defaults(revision: 'patch')
  unless AVAILABLE_REVISIONS.include?(args.revision)
    abort "Please provide valid revision: #{AVAILABLE_REVISIONS.join(',')}"
  end

  system "bumpversion #{args.revision}"
end
# -----------------------------------------------------------------------------


# default task
# -----------------------------------------------------------------------------
desc 'show avaliable tasks (default task)'
task :default do
  system('rake -sT')
end
# -----------------------------------------------------------------------------


# release new version
# -----------------------------------------------------------------------------
desc "release new version #{AVAILABLE_REVISIONS.join(',')}, default: patch"
task :release, [:revision] => [:repo_clean] do |_, args|
  args.with_defaults(revision: 'patch')
  Rake::Task['bump'].invoke(args.revision)
end
# -----------------------------------------------------------------------------
