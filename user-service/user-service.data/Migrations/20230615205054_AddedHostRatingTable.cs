using System;
using Microsoft.EntityFrameworkCore.Migrations;

#nullable disable

namespace user_service.data.Migrations
{
    /// <inheritdoc />
    public partial class AddedHostRatingTable : Migration
    {
        /// <inheritdoc />
        protected override void Up(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.CreateTable(
                name: "HostRatings",
                columns: table => new
                {
                    Id = table.Column<Guid>(type: "char(36)", nullable: false),
                    AverageRating = table.Column<double>(type: "double", nullable: false),
                    NumOfRatings = table.Column<uint>(type: "int unsigned", nullable: false),
                    TotalRating = table.Column<uint>(type: "int unsigned", nullable: false),
                    HostId = table.Column<Guid>(type: "char(36)", nullable: false)
                },
                constraints: table =>
                {
                    table.PrimaryKey("PK_HostRatings", x => x.Id);
                    table.ForeignKey(
                        name: "FK_HostRatings_Users_HostId",
                        column: x => x.HostId,
                        principalTable: "Users",
                        principalColumn: "Id",
                        onDelete: ReferentialAction.Cascade);
                })
                .Annotation("MySQL:Charset", "utf8mb4");

            migrationBuilder.CreateIndex(
                name: "IX_HostRatings_HostId",
                table: "HostRatings",
                column: "HostId",
                unique: true);
        }

        /// <inheritdoc />
        protected override void Down(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.DropTable(
                name: "HostRatings");
        }
    }
}
