/*
 Navicat Premium Data Transfer

 Source Server         : 127.0.0.1
 Source Server Type    : MySQL
 Source Server Version : 50726
 Source Host           : 127.0.0.1:3306
 Source Schema         : CCB

 Target Server Type    : MySQL
 Target Server Version : 50726
 File Encoding         : 65001

 Date: 30/08/2019 11:57:47
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for menu
-- ----------------------------
DROP TABLE IF EXISTS `menu`;
CREATE TABLE `menu`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `parent_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT 'id do pai',
  `type` tinyint(3) UNSIGNED NOT NULL DEFAULT 1 COMMENT 'Tipo de menu;1:Existe uma interface para acessar o menu,2:Nenhum menu de interface acessível,3:Apenas como um menu',
  `status` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT 'Estado;1:exposição,0:Não mostre',
  `list_order` tinyint(3) UNSIGNED NOT NULL DEFAULT 100 COMMENT 'Ordenação',
  `controller` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT 'Nome do controlador',
  `action` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT 'Nome da Operação',
  `param` varchar(250) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT 'Parâmetros extras',
  `name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT 'Nome do menu',
  `icon` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT 'Ícone do menu',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT 'Observação',
  `module_belong` tinyint(3) NOT NULL DEFAULT 1 COMMENT 'Leia a configuração do módulo',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `status`(`status`) USING BTREE,
  INDEX `parent_id`(`parent_id`) USING BTREE,
  INDEX `controller`(`controller`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 720 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = 'Mesa de menu de bastidores' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of menu
-- ----------------------------
INSERT INTO `menu` VALUES (1, 0, 1, 1, 2, 'IndexController', 'Main', '', 'home', '', '', 1);
INSERT INTO `menu` VALUES (6, 0, 0, 1, 100, 'Setting', 'default', '', 'configurar', 'layui-icon-setting', 'Entrada de configurações do sistema', 1);
INSERT INTO `menu` VALUES (20, 0, 1, 1, 100, 'MenuController', 'default', '', 'Menu', 'layui-icon-menu', 'Gerenciamento de menu', 1);
INSERT INTO `menu` VALUES (21, 20, 1, 1, 100, 'MenuController', 'Lists', '', 'Todos', 'layui-icon-menu', 'Lista de todos os menus', 1);
INSERT INTO `menu` VALUES (22, 20, 1, 1, 100, 'MenuController', 'Save', '', 'Menu adicionado', '', 'Menu adicionado', 1);
INSERT INTO `menu` VALUES (23, 20, 2, 1, 100, 'Menu', 'addPost', '', 'Adicione no menu', '', 'Adicione no menu', 1);
INSERT INTO `menu` VALUES (24, 20, 3, 1, 100, 'Menu', 'edit', '', 'Edição do menu', '', 'Edição do menu', 1);
INSERT INTO `menu` VALUES (25, 20, 2, 1, 100, 'Menu', 'editPost', '', 'Editar Menu, enviar e salvar', '', 'Menu editar, enviar e salvar', 1);
INSERT INTO `menu` VALUES (26, 20, 2, 1, 100, 'Menu', 'delete', '', 'Excluir menu', '', 'Excluir menu', 1);
INSERT INTO `menu` VALUES (27, 20, 2, 1, 100, 'Menu', 'listOrder', '', 'Classificação ', '', 'Classificação', 1);
INSERT INTO `menu` VALUES (49, 109, 0, 1, 100, 'User', 'default', '', 'Grupo de Gestores', '', '管理组', 1);
INSERT INTO `menu` VALUES (50, 49, 1, 1, 100, 'Rbac', 'index', 'type=1', 'Gestão de funções', '', '角色管理', 1);
INSERT INTO `menu` VALUES (51, 50, 1, 1, 100, 'Rbac', 'roleAdd', '', 'Adicionar papel', '', '添加角色', 1);
INSERT INTO `menu` VALUES (52, 50, 2, 1, 100, 'Rbac', 'roleAddPost', '', 'Adicionar envio de função', '', 'Adicionar envio de função', 1);
INSERT INTO `menu` VALUES (53, 50, 1, 1, 100, 'Rbac', 'roleEdit', '', 'Editar papel', '', 'Editar papel', 1);
INSERT INTO `menu` VALUES (54, 50, 2, 1, 100, 'Rbac', 'roleEditPost', '', 'Editar envio de função', '', 'Editar envio de função', 1);
INSERT INTO `menu` VALUES (55, 50, 2, 1, 100, 'Rbac', 'roleDelete', '', 'Excluir papel', '', 'Excluir papel', 1);
INSERT INTO `menu` VALUES (56, 50, 1, 1, 100, 'Rbac', 'authorize', '', 'Definir permissões de função', '', 'Definir permissões de função', 1);
INSERT INTO `menu` VALUES (57, 50, 2, 1, 100, 'Rbac', 'authorizePost', '', 'Envio de autorização de função', '', 'Envio de autorização de função', 1);
INSERT INTO `menu` VALUES (109, 0, 0, 1, 100, 'AdminIndex', 'default', '', 'Gestão de Usuários', 'group', 'Gestão de Usuários', 1);
INSERT INTO `menu` VALUES (110, 49, 1, 1, 100, 'User', 'Index', 'type=1', 'administrador', '', 'Gestão de administrador', 1);
INSERT INTO `menu` VALUES (111, 110, 1, 1, 100, 'User', 'add', '', 'Novo usuário', '', 'Inserir um novo usuário', 1);
INSERT INTO `menu` VALUES (112, 111, 1, 1, 100, 'User', 'add', '', 'Novo usuário', '', '管理员添加', 1);
INSERT INTO `menu` VALUES (113, 111, 2, 1, 100, 'User', 'addPost', '', 'Admin adicionar submissão', '', 'Admin adicionar submissão', 1);
INSERT INTO `menu` VALUES (114, 111, 1, 1, 100, 'User', 'edit', '', 'Edição de administrador', '', 'Edição de administrador', 1);
INSERT INTO `menu` VALUES (115, 111, 2, 1, 100, 'User', 'editPost', '', 'edição de administrador', '', 'edição de administrador', 1);
INSERT INTO `menu` VALUES (116, 111, 1, 1, 100, 'User', 'userInfo', '', 'Informação pessoal', '', 'Modificação de informações pessoais do administrador', 1);
INSERT INTO `menu` VALUES (117, 111, 2, 1, 100, 'User', 'userInfoPost', '', 'Modif de informações pessoais', '', 'Modif de informações pessoais', 1);
INSERT INTO `menu` VALUES (118, 111, 2, 1, 100, 'User', 'delete', '', 'Exclusão de administrador', '', 'Exclusão de administrador', 1);
INSERT INTO `menu` VALUES (119, 111, 2, 1, 100, 'User', 'ban', '', 'Desativar administrador', '', 'Desativar administrador', 1);
INSERT INTO `menu` VALUES (120, 111, 2, 1, 100, 'User', 'cancelBan', '', 'Habilitar administrador', '', 'Habilitar administrador', 1);
INSERT INTO `menu` VALUES (711, 0, 1, 1, 1, 'SysController', 'Config', '', 'Configuração do site', 'layui-icon-control', 'Configure algumas informações do site, avatar, logotipo, etc.', 1);
INSERT INTO `menu` VALUES (720, 0, 0, 1, 4, 'MenuController', 'default', '', 'Cadastros', 'layui-icon-menu', '', 1);
INSERT INTO `menu` VALUES (721, 720, 1, 1, 1, 'ClientController', 'Lists', '', 'Cliente', '', '', 1);

-- ----------------------------
-- Table structure for option
-- ----------------------------
DROP TABLE IF EXISTS `option`;
CREATE TABLE `option`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `autoload` tinyint(3) UNSIGNED NOT NULL DEFAULT 1 COMMENT 'Para carregar automaticamente;1:Carregamento automático;0:Não carregue automaticamente',
  `option_name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'Nome da configuração',
  `option_value` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL COMMENT 'Valor de configuração',
  `option_comment` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'Comentários chineses',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `option_name`(`option_name`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 60 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = 'Tabela de configuração de todo o site' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of option
-- ----------------------------
INSERT INTO `option` VALUES (59, 0, 'site_config', '{\"site_logo\":\"upload/pic/2019/08/30/3931623433393835663033636531323166363762616439343235303533643535.png\",\"site_name\":\"pm\",\"remark\":\"CCB是基于Beego开发的易用、易扩展、界面友好的轻量级功能权限管理系统。\\n前端框架基于layui进行资源整合。\\n本系统基于beego开发，默认使用mysql数据库，缓存redis \"}', '');

-- ----------------------------
-- Table structure for user
-- ----------------------------
-- DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `true_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '',
  `login` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `pwd` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `role` tinyint(1) NULL DEFAULT NULL COMMENT '1默认管理员',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `name`(`true_name`) USING BTREE,
  UNIQUE INDEX `email`(`login`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 21 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = 'mesa do usuário' ROW_FORMAT = Dynamic;


-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (7, 'Administrador', 'admin', 'upload/pic/2019/08/09/d6c4821b8532cda9c43c67a91a614e4f.jpg', '82790085228cf8a1e3bac41f45271e5f', 1);

SET FOREIGN_KEY_CHECKS = 1;


-- ----------------------------
-- Table structure for clientes
-- ----------------------------
-- DROP TABLE IF EXISTS `clientes`;
CREATE TABLE `clientes` (
	`id` INT(5) UNSIGNED NOT NULL AUTO_INCREMENT,
	`nome` VARCHAR(120) NOT NULL COLLATE 'latin1_swedish_ci',
	`contato_funcao` VARCHAR(20) NULL DEFAULT NULL COLLATE 'latin1_swedish_ci',
	`contato_nome` VARCHAR(40) NULL DEFAULT NULL COLLATE 'latin1_swedish_ci',
	`cgc_cpf` VARCHAR(20) NULL DEFAULT NULL COLLATE 'latin1_swedish_ci',
	`Razao_social` VARCHAR(100) NULL DEFAULT NULL COLLATE 'utf8_general_ci',
	`inscr_estadual` VARCHAR(20) NULL DEFAULT NULL COLLATE 'latin1_swedish_ci',
	`endereco` VARCHAR(100) NULL DEFAULT NULL COLLATE 'latin1_swedish_ci',
	`cidade` VARCHAR(40) NOT NULL COLLATE 'latin1_swedish_ci',
	`uf` VARCHAR(2) NULL DEFAULT NULL COLLATE 'latin1_swedish_ci',
	`cep` VARCHAR(9) NULL DEFAULT NULL COLLATE 'latin1_swedish_ci',
	`telefone1` VARCHAR(20) NULL DEFAULT NULL COLLATE 'latin1_swedish_ci',
	`telefone2` VARCHAR(20) NULL DEFAULT NULL COLLATE 'latin1_swedish_ci',
	`telefone3` VARCHAR(20) NULL DEFAULT NULL COLLATE 'latin1_swedish_ci',
	`email` VARCHAR(60) NULL DEFAULT NULL COLLATE 'latin1_swedish_ci',
	`obs` VARCHAR(1024) NULL DEFAULT NULL COLLATE 'latin1_swedish_ci',
	`estado` TINYINT(4) NULL DEFAULT '1',
	`preco_base` DECIMAL(6,2) NULL DEFAULT NULL,
	PRIMARY KEY (`id`) USING BTREE
)
COLLATE='utf8_general_ci'
ENGINE=InnoDB
ROW_FORMAT=DYNAMIC
AUTO_INCREMENT=5810
;


