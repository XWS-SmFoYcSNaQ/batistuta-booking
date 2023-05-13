using Microsoft.EntityFrameworkCore;
using user_service.domain.Entities;
using user_service.domain.Enums;

namespace user_service.data.Db
{
    public class UserServiceDbContext : DbContext
    {
        public DbSet<User> Users { get; set; }

        public UserServiceDbContext(DbContextOptions options) : base(options) { }

        protected override void OnModelCreating(ModelBuilder modelBuilder)
        {
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

            // Seed users
            modelBuilder.Entity<User>().HasData(new[]
            {
                new User()
                {
                    Id = Guid.Parse("00695ec1-8ed0-4781-a432-9fdd98bd6cde"),
                    Role = UserRole.Nonauthenticated,
                    Email = "nonauthenticated@test.com",
                    Username = "nonauthenticated",
                    Password = "AMhtGKXzeetCyumUykEQ2R0YxxLzYTJPzeQbaLQeseWBqf8bQiEmcUQV/0XRUkvQUA==",
                    FirstName = "Non",
                    LastName = "Authenticated",
                    LivingPlace = "Beograd, Srbija"
                },
                new User()
                {
                    Id = Guid.Parse("422bd1e1-6d70-42cc-adc6-89a09b313c01"),
                    Role = UserRole.Guest,
                    Email = "guest@test.com",
                    Username = "guest",
                    Password = "AHD4mlEmnCAhyWDp4D6H3lcLzr5WVROCnshsLiuitVapL7GWNshdiY4pFElnRqN0qQ==",
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
                    Password = "AEYedJWK/IHAD3NNv03runGcGLeN1CDKyKZS8ni/3x3gFKP8AwN5m+F2vX2kH/McjQ==",
                    FirstName = "Host",
                    LastName = "Host",
                    LivingPlace = "Sabac, Srbija"
                }
            });
        }
    }
}
