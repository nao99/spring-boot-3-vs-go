package org.ndbs.spring.graalvm.product.api;

import org.ndbs.spring.graalvm.product.api.model.CreateProductDto;
import org.ndbs.spring.graalvm.product.domain.ProductService;
import org.ndbs.spring.graalvm.product.domain.model.Product;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.web.bind.annotation.*;
import org.springframework.web.server.ResponseStatusException;

import java.util.Set;

/**
 * ProductController class <br>
 *
 * <br>
 *
 * @author Nikolai Osipov <nao99.dev@gmail.com>
 * @since  2022-12-11
 */
@RestController
@RequestMapping("/api/v1/products")
public class ProductController {
    private final ProductService service;

    @Autowired
    public ProductController(ProductService service) {
        this.service = service;
    }

    @GetMapping("/{productId}")
    public Product getProduct(@PathVariable("productId") Long id) {
        try {
            return service.getProductById(id);
        } catch (RuntimeException e) {
            throw new ResponseStatusException(HttpStatus.NOT_FOUND);
        }
    }

    @GetMapping
    public Set<Product> getProducts() {
        return service.getProducts();
    }

    @PostMapping
    public Product createProduct(@RequestBody CreateProductDto createDto) {
        return service.createProduct(createDto);
    }
}
