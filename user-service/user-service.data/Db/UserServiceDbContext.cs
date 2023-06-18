using Microsoft.EntityFrameworkCore;
using user_service.domain.Entities;
using user_service.domain.Enums;

namespace user_service.data.Db
{
    public class UserServiceDbContext : DbContext
    {
        public DbSet<User> Users { get; set; }
        public DbSet<HostRating> HostRatings { get; set; }

        public UserServiceDbContext(DbContextOptions options) : base(options) { }

        protected override void OnModelCreating(ModelBuilder modelBuilder)
        {
            modelBuilder.Entity<User>()
                .Property(x => x.Id)
                .ValueGeneratedOnAdd();
            modelBuilder.Entity<User>().HasAlternateKey(x => x.Username);
            modelBuilder.Entity<User>().HasAlternateKey(x => x.Email);
            modelBuilder.Entity<User>()
                .Property(x => x.Role)
                .HasConversion(
                    x => x.ToString(),
                    x => (UserRole)Enum.Parse(typeof(UserRole), x))
                .HasMaxLength(32);

            modelBuilder.Entity<User>().Property(x => x.Username).HasMaxLength(64);
            modelBuilder.Entity<User>().Property(x => x.Password).HasMaxLength(256);
            modelBuilder.Entity<User>().Property(x => x.Email).HasMaxLength(64);
            modelBuilder.Entity<User>().Property(x => x.FirstName).HasMaxLength(32);
            modelBuilder.Entity<User>().Property(x => x.LastName).HasMaxLength(32);
            modelBuilder.Entity<User>().Property(x => x.LivingPlace).HasMaxLength(64);

            modelBuilder.Entity<HostRating>()
                .Property(x => x.Id)
                .ValueGeneratedOnAdd();
            modelBuilder.Entity<HostRating>()
                .HasOne(x => x.Host)
                .WithOne()
                .HasForeignKey<HostRating>(x => x.HostId)
                .IsRequired(true);

            // Seed users
            modelBuilder.Entity<User>().HasData(new[]
            {
                new User()
                {
                    Id = Guid.Parse("422bd1e1-6d70-42cc-adc6-89a09b313c01"),
                    Role = UserRole.Guest,
                    Email = "guest@test.com",
                    Username = "guest",
                    Password = "10000.GhMJYLVMJDUSKPYAt3G+oA==.2d2SyAT1CWcY/eNqJiKXKdTvjrWY2TftfJsHiOCy54g=",
                    FirstName = "Guest",
                    LastName = "Guest",
                    LivingPlace = "Novi Sad, Srbija"
                },
                new User()
                {
                    Id = Guid.Parse("87abe34c-9935-44c4-aad5-af82f9442c77"),
                    Role = UserRole.Host,
                    Email = "host@test.com",
                    Username = "host",
                    Password = "10000.pUrW8b1z1nt7+RFCVYWpWg==.jp+r7PJ49rgJwwZVAfLhCb2YMyCAZR3gXgrLnson2UQ=",
                    FirstName = "Host",
                    LastName = "Host",
                    LivingPlace = "Sabac, Srbija"
                }
            });
        }
    }
}
